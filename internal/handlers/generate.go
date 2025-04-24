package handlers

import (
	"context"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func GenerateCode(ctx context.Context, msgs []types.Message) (string, error) {

	backend, err := agent.GetBackend(config.Config)
	if err != nil {
		msg := "Error getting backend: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	resp, err := backend.Respond(ctx, msgs)
	if err != nil {
		msg := "Error getting response: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	return resp.GetResponse(), nil
}
func GenerateCodePrompt(prompt string, isFileOp bool) []types.Message {
	msgs := []types.Message{
		{
			Role:    "system",
			Content: "You are a highly skilled software developer. Your task is to generate concise and correct code snippets in response to user requests. You are capable of writing code in languages such as Go, Python, JavaScript, and C++, as well as IaC tools like Terraform and Docker.",
		},
		{
			Role: "system",
			Content: "Always follow these rules strictly:\n" +
				"1. Never explain or describe the code.\n" +
				"2. Never include headings or extra text.\n" +
				"3. Only return structured JSON or plain code output as instructed.\n" +
				"4. Try to understand the problem and write the code accordingly.\n" +
				"5. Always five full code as if you have to write the whole file yourself.\n",
		},
	}

	if isFileOp {
		msgs = append(msgs,
			types.Message{Role: "system", Content: `When generating file-based output, respond only in this format:
{
  "file_name": "main.go",
  "code_snippet": "...."
}`},
		)
	} else {
		msgs = append(msgs,
			types.Message{Role: "system", Content: `Respond with code only, NO EXPLANATIONS , and DO NOT wrap it in triple backticks with the correct language identifier.`},
		)
	}

	msgs = append(msgs, types.Message{
		Role:    "user",
		Content: prompt,
	})

	return msgs
}
