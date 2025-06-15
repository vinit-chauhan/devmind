package openai

import (
	"context"
	"strings"

	"github.com/openai/openai-go"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/consumer"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var (
	tknCh  = make(chan string, 16)
	doneCh = make(chan struct{}, 1)
	errCh  = make(chan error, 1)
)

func (b *OpenAIBackend) Respond(ctx context.Context, msgs []types.Message) (types.Readable, error) {
	full := strings.Builder{}
	defer func() {
		close(errCh)
	}()

	// Produce the response in a goroutine
	go Produce(ctx, b, msgs, &full)

	// Print the response as it comes to stdout
	go consumer.Consume(ctx, tknCh, doneCh)

	// Wait for the response to finish
	<-doneCh

	parsed := &OpenAIResponse{
		Response: types.Message{
			Role:    "assistant",
			Content: full.String(),
		},
		Done: true,
	}

	return parsed, nil
}

func Produce(ctx context.Context, b *OpenAIBackend, msgs []types.Message, full *strings.Builder) {
	defer func() {
		logger.Debug("Closing OpenAI Producer")
		close(tknCh)
	}()

	messages := make([]openai.ChatCompletionMessageParamUnion, 0)
	for _, msg := range msgs {
		switch msg.Role {
		case string(openai.MessageRoleUser):
			messages = append(messages, openai.UserMessage(msg.Content))
		case string(openai.MessageRoleAssistant):
			messages = append(messages, openai.AssistantMessage(msg.Content))
		}
	}

	stream := b.client.Chat.Completions.NewStreaming(
		ctx,
		openai.ChatCompletionNewParams{
			Messages: messages,
			Model:    b.conf.Model,
		},
	)

	for stream.Next() {
		curr := stream.Current().Choices[0]
		if content := curr.Delta.Content; content != "" {
			select {
			case <-ctx.Done():
				logger.Debug("Context done in OpenAI chat")
				return
			case tknCh <- content:
				full.WriteString(content)
			}
		}
	}

	if err := stream.Err(); err != nil {
		logger.Error(err.Error())
		errCh <- err
	}
}
