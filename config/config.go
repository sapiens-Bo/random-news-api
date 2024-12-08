package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API       string
	TgAPI     string
	TgBotHost string
}

func MustLoad() *Config {
	const op = "config.MustLoad"
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err config file load: %v", err)
	}
	cfg.API = os.Getenv("API_NEWS")
	cfg.TgAPI = os.Getenv("TG_BOT_API")
	cfg.TgBotHost = "api.telegram.org"

	return &cfg
}
