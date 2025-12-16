package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type StealthConfig struct {
	MinThinkMs int `yaml:"min_think_ms"`
	MaxThinkMs int `yaml:"max_think_ms"`
}

type BrowserConfig struct {
	Headless bool `yaml:"headless"`
}

type Config struct {
	Browser BrowserConfig `yaml:"browser"`
	Stealth StealthConfig `yaml:"stealth"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{
		Browser: BrowserConfig{
			Headless: false,
		},
		Stealth: StealthConfig{
			MinThinkMs: 800,
			MaxThinkMs: 1500,
		},
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, nil // defaults if file missing
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
