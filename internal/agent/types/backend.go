package types

type Backend interface {
	Respond(prompt string) (Readable, error)
}
