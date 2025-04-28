package chat

import (
	"fmt"
	"strings"
	"time"
)

var ChatParseErr = fmt.Errorf("chat parse error")

type Chat struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func (c *Chat) String() string {
	return fmt.Sprintf("%s %s: %s", c.CreatedAt, c.Role, c.Content)
}

func (c *Chat) Parse(chat string) error {
	chat = strings.TrimSpace(chat)
	parts := strings.SplitN(chat, " ", 3)
	if len(parts) < 3 {
		return ChatParseErr
	}
	c.CreatedAt = parts[0]
	c.Role = parts[1]
	c.Content = parts[2]

	return nil
}

func New(role, content string) Chat {
	return Chat{
		Role:      role,
		Content:   content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
