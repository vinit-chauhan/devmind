package config

var Config AppConfig

type OpenAIConfig struct {
	ApiKey string `mapstructure:"api_key,omitempty"`
	Model  string `mapstructure:"model,omitempty"`
}

type OllamaConfig struct {
	Model  string `mapstructure:"model,omitempty"`
	Host   string `mapstructure:"host,omitempty"`
	Stream bool   `mapstructure:"stream,omitempty"`
}

type AppConfig struct {
	Backend string       `mapstructure:"backend,omitempty"`
	OpenAI  OpenAIConfig `mapstructure:"openai,omitempty"`
	Ollama  OllamaConfig `mapstructure:"ollama,omitempty"`
}
