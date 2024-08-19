package config

import (
	"errors"
	"os"
)

type Config struct {
	Port string
}

func LoadFromEnv() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return Config{}, errors.New("PORT is not found in the enviroment")
	}

	cfg := Config{}
	cfg.Port = port

	return cfg, nil
}
