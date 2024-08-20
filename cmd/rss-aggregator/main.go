package main

import (
	"log"

	"github.com/5aradise/rss-aggregator/config"
	"github.com/5aradise/rss-aggregator/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
