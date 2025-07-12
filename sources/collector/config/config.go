package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string `env:"ENV"`
	LogLevel string `env:"LOG_LEVEL"`
	XApiKey  string `env:"X_API_KEY,notEmpty"`
}

func Load() (*Config, error) {
	// Load .env, if present
	godotenv.Load()

	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
