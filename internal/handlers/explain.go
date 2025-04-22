package handlers

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

func Explain(ctx context.Context, filename string, lr utils.LineRange) (string, error) {
	prompt, err := generatePrompt(filename, lr)
	if err != nil {
		msg := "Error generating prompt: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	backend, err := agent.GetBackend(config.Config)
	if err != nil {
		msg := "Error getting backend: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	resp, err := backend.Respond(ctx, prompt)
	if err != nil {
		msg := "Error getting response: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	return resp.GetResponse(), nil
}

func generatePrompt(filename string, lr utils.LineRange) (string, error) {
	logger.Debug(fmt.Sprintf("Explaining lines %d-%d of file %s", lr.Start, lr.End, filename))
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("Error opening file: %s", err.Error())

	}
	defer file.Close()

	logger.Debug("Reading content of file " + filename)
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %s", err.Error())
	}

	if lr.IsValid() {
		logger.Debug("Extracting lines " + lr.String() + " from file " + filename)
		extractedContent, err := lr.ExtractLines(string(content))
		if err != nil {
			return "", fmt.Errorf("Error extracting lines: %s", err.Error())
		}
		content = []byte(extractedContent)
	}

	prompt := strings.Builder{}
	prompt.WriteString("Explain the following code snippet in detail:\n\n")
	prompt.WriteString("```go\n")
	prompt.Write(content)
	prompt.WriteString("```\n\n")
	prompt.WriteString("Only provide a detailed explanation of the code snippet above.")

	return prompt.String(), nil
}
