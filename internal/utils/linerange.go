package utils

import (
	"fmt"
	"strings"

	"github.com/vinit-chauhan/devmind/internal/logger"
)

type LineRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func ParseLineRange(lineRange string) (LineRange, error) {
	var lr LineRange

	if lineRange == "" {
		return LineRange{}, nil
	}

	_, err := fmt.Sscanf(lineRange, "%d-%d", &lr.Start, &lr.End)
	if err != nil {
		return LineRange{}, err
	}
	return lr, nil
}

func (lr LineRange) IsValid() bool {
	return lr.Start > 0 && lr.End > 0 && lr.Start <= lr.End
}

func (lr LineRange) String() string {
	if lr.IsValid() {
		return fmt.Sprintf("%d-%d", lr.Start, lr.End)
	}
	return ""
}

func (lr LineRange) ExtractLines(content string) (string, error) {
	lines := strings.Split(content, "\n")
	if lr.Start > len(lines) {
		return "", fmt.Errorf("line range out of bounds")
	}

	if lr.End > len(lines) {
		logger.Debug("End line exceeds total lines, adjusting to total lines: " + fmt.Sprint(len(lines)))
		lr.End = len(lines)
	}
	extractedLines := strings.Join(lines[lr.Start-1:lr.End], "\n")
	return extractedLines, nil
}
