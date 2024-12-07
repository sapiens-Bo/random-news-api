package config

import "os"

type Config struct {
	API string
}

func MustLoad() *Config {
	var cfg Config

	cfg.API = os.Getenv("API_NEWS")

	return &cfg
}
