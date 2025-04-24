package types

import (
	"context"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Backend interface {
	Respond(ctx context.Context, prompt []Message) (Readable, error)
}
