package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("Failed to parse YAML configuration file: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, fmt.Errorf("input-file is missing")
	}

	if cfg.OutputFile == "" {
		return nil, fmt.Errorf("output-file is missing")
	}

	return &cfg, nil
}
