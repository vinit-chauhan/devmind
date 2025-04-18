package config

import (
	"github.com/spf13/viper"

	"github.com/vinit-chauhan/devmind/internal/logger"
)

var Config AppConfig

type OpenAIConfig struct {
	APIKey string `mapstructure:"api_key,omitempty"`
	APIURL string `mapstructure:"api_url,omitempty"`
}

type OLLAMAConfig struct {
	Model string `mapstructure:"model,omitempty"`
	Host  string `mapstructure:"host,omitempty"`
}

type AppConfig struct {
	Backend string       `mapstructure:"backend,omitempty"`
	OpenAI  OpenAIConfig `mapstructure:"openai,omitempty"`
	Ollama  OLLAMAConfig `mapstructure:"ollama,omitempty"`
}

func InitConfig() {
	viper.SetConfigName("devmind")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.devmind")
	viper.AddConfigPath("/etc/devmind")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading config file: " + err.Error())
		return
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.Error("Error unmarshalling config file: " + err.Error())
		return
	}

	logger.Info("Config file loaded successfully")
}
