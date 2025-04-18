package config

import (
	"github.com/spf13/viper"

	"github.com/vinit-chauhan/devmind/internal/logger"
)

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
