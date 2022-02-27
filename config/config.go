package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Conn int
	Urls []string
}

func New() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Conn: viper.GetInt("conn"),
		Urls: viper.GetStringSlice("urls"),
	}, nil
}
