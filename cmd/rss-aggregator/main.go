package main

import (
	"log"
	"path/filepath"

	"github.com/5aradise/rss-aggregator/config"
	"github.com/5aradise/rss-aggregator/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(filepath.Join("..", "..", ".env"))

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(config.Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
