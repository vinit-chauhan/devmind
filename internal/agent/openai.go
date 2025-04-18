package agent

type OpenAIBackend struct{}

func NewOpenAIBackend() Backend {
	return &OpenAIBackend{}
}

func (b *OpenAIBackend) Respond(prompt string) (response string, err error) {
	return "OpenAI response", nil
}
