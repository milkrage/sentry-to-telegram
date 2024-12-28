package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Sentry Sentry `yaml:"sentry"`
}

type Sentry struct {
	Hostname     string `yaml:"hostname"      env:"SENTRY_HOSTNAME"`
	ClientID     string `yaml:"client_id"     env:"SENTRY_CLIENT_ID"`
	ClientSecret string `yaml:"client_secret" env:"SENTRY_CLIENT_SECRET"`
}

func New(path string) (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	err = env.Parse(config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env variables: %w", err)
	}

	return config, nil
}
