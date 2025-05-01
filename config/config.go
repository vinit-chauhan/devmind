package config

import (
	"runtime"

	"github.com/spf13/viper"

	"github.com/vinit-chauhan/devmind/internal/logger"
)

func InitConfig() {
	viper.SetConfigName("devmind")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	if runtime.GOOS == "windows" {
		viper.AddConfigPath("%APPDATA%/devmind")
		viper.AddConfigPath("%LOCALAPPDATA%/devmind")
		viper.AddConfigPath("%PROGRAMDATA%/devmind")
	} else if runtime.GOOS == "linux" {
		viper.AddConfigPath("$XDG_CONFIG_HOME/devmind")
		viper.AddConfigPath("$HOME/.devmind")
		viper.AddConfigPath("/etc/devmind")
	} else if runtime.GOOS == "darwin" {
		viper.AddConfigPath("$HOME/Library/Application Support/devmind")
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
