package config

import (
	"errors"
	"os"
)

type Config struct {
	Port  string
	DbURL string
}

func LoadFromEnv() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return Config{}, errors.New("PORT is not found in the enviroment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return Config{}, errors.New("DB_URL is not found in the enviroment")
	}

	cfg := Config{}
	cfg.Port = port
	cfg.DbURL = dbURL

	return cfg, nil
}
