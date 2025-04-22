package types

type Backend interface {
	Respond(prompt string) (Response, error)
}
