package configs

import (
	"errors"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string         `env:"ENV"`
	Port     string         `env:"PORT"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
}

type PostgresConfig struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
}

func NewConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load .env file: " + err.Error())
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config from environment variables: " + err.Error())
	}

	return cfg, nil
}
