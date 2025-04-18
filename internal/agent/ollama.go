package agent

type OllamaBackend struct {
}

func NewOllamaBackend() Backend {
	return &OllamaBackend{}
}

func (b *OllamaBackend) Respond(prompt string) (response string, err error) {
	return "Ollama response", nil
}
