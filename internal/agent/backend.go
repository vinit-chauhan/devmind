package agent

type Backend interface {
	Respond(prompt string) (string, error)
}
