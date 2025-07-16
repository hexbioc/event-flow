package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string `env:"ENV"`
	LogLevel string `env:"LOG_LEVEL"`
	XApiKey  string `env:"X_API_KEY,notEmpty"`

	RMQTLS      string `env:"RMQ_TLS"`
	RMQHostname string `env:"RMQ_HOSTNAME"`
	RMQUser     string `env:"RMQ_USER"`
	RMQPassword string `env:"RMQ_PASSWORD"`
	RMQVhost    string `env:"RMQ_VHOST"`
}

func Load() (*Config, error) {
	// Load .env, if present
	godotenv.Load()

	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
