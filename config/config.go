package config

import (
	"github.com/spf13/viper"

	"github.com/vinit-chauhan/devmind/internal/constants"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func Init() {
	viper.SetConfigName("devmind")
	viper.SetConfigType("yaml")

	for _, path := range constants.CONFIG_PATHS {
		viper.AddConfigPath(path)
	}

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
