package config

import "github.com/spf13/viper"

type Config struct {
	Port     string `mapstructure:"SERVER_PORT"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "DEBUG")

	var cfg Config
	err := viper.Unmarshal(&cfg)

	return &cfg, err
}
