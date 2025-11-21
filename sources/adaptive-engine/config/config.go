package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	// HTTP Server Configuration
	Port int    `env:"APP_PORT" envDefault:"8084"`
	Mode string `env:"API_MODE" envDefault:"debug"`

	// External Services
	LearnerServiceURL string `env:"LEARNER_SERVICE_URL" envDefault:"http://localhost:8083"`
	ContentServiceURL string `env:"CONTENT_SERVICE_URL" envDefault:"http://localhost:8081"`

	// Logger Configuration
	LoggerLevel    string `env:"LOGGER_LEVEL" envDefault:"info"`
	LoggerMode     string `env:"LOGGER_MODE" envDefault:"debug"`
	LoggerEncoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

// Load is the function to load the configuration from the environment variables.
func Load() (*Config, error) {
	// Load .env file if exists (ignore error if not exists)
	_ = godotenv.Load()

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
