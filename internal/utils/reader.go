package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/vinit-chauhan/devmind/internal/logger"
)

func ReadFileContent(filename string, lr LineRange) ([]byte, error) {
	logger.Debug(fmt.Sprintf("Explaining lines %d-%d of file %s", lr.Start, lr.End, filename))

	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err.Error())
	}
	defer file.Close()

	logger.Debug("Reading content of file " + filename)
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %s", err.Error())
	}

	if lr.IsValid() {
		logger.Debug("Extracting lines " + lr.String() + " from file " + filename)
		extractedContent, err := lr.ExtractLines(string(content))
		if err != nil {
			return nil, fmt.Errorf("Error extracting lines: %s", err.Error())
		}
		content = []byte(extractedContent)
	}

	return content, nil
}

func ReadStdin() ([]byte, error) {
	// Read from stdin until EOF
	var data []byte
	for {
		buf := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		data = append(data, buf[:n]...)
	}
	return data, nil
}
