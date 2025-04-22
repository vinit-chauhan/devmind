package types

import "context"

type Backend interface {
	Respond(ctx context.Context, prompt string) (Readable, error)
}
