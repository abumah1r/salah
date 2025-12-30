package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.salah")

	// Don't error if config file doesn't exist yet
	viper.ReadInConfig()

	return nil
}
