package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		Http     `yaml:"http"`
		Postgres `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME"         env-required:"true" yaml:"name"`
		Version string `env:"APP_VERSION"      env-required:"true" yaml:"version"`
	}

	// Http -.
	Http struct {
		Port string `env:"HTTP_PORT" env-required:"true" yaml:"port"`
	}

	// Postgres -.
	Postgres struct {
		Host     string `env:"POSTGRES_HOST"     env-required:"true" yaml:"host"`
		Port     int    `env:"POSTGRES_PORT"     env-required:"true" yaml:"port"`
		User     string `env:"POSTGRES_USER"     env-required:"true" yaml:"user"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true" yaml:"password"`
		DBName   string `env:"POSTGRES_DB_NAME"  env-required:"true" yaml:"dbname"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	if configPath != "" {
		err := cleanenv.ReadConfig(configPath, cfg)
		if err != nil {
			return nil, errors.New("getting configuration from file: " + err.Error())
		}
	}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, errors.New("getting configuration from environment variables: " + err.Error())
	}
	return cfg, nil
}
