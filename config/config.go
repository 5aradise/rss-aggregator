package config

import (
	"errors"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

var Cfg Config

func Load() error {
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("PORT is not found in the enviroment")
	}

	Cfg.Server.Port = port

	return nil
}
