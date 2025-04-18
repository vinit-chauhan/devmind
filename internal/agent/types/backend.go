package types

type Backend interface {
	Respond(prompt string) (string, error)
}
