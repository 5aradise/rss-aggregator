package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server struct {
		Port string
	}
	DB struct {
		URL string
	}
	RSS struct {
		ConurentRequests    int32
		TimeToRequest       time.Duration
		TimeBetweenRequests time.Duration
	}
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

	conurentRequestsStr := os.Getenv("REQUESTS")
	if conurentRequestsStr == "" {
		return Config{}, errors.New("REQUESTS is not found in the enviroment")
	}
	conurentRequests, err := strconv.ParseInt(conurentRequestsStr, 10, 32)
	if err != nil {
		return Config{}, err
	}

	strStr := os.Getenv("STR")
	if strStr == "" {
		return Config{}, errors.New("STR is not found in the enviroment")
	}
	str, err := strconv.Atoi(strStr)
	if err != nil {
		return Config{}, err
	}

	sbrStr := os.Getenv("SBR")
	if strStr == "" {
		return Config{}, errors.New("SBR is not found in the enviroment")
	}
	sbr, err := strconv.Atoi(sbrStr)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	cfg.Server.Port = port

	cfg.DB.URL = dbURL

	cfg.RSS.ConurentRequests = int32(conurentRequests)
	cfg.RSS.TimeToRequest = time.Duration(str) * time.Second
	cfg.RSS.TimeBetweenRequests = time.Duration(sbr) * time.Second

	return cfg, nil
}
