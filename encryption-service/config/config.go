package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
