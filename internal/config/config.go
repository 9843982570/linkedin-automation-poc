package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Browser struct {
		Headless     bool `yaml:"headless"`
		SlowMotionMs int  `yaml:"slow_motion_ms"`
	} `yaml:"browser"`

	Limits struct {
		DailyConnections int `yaml:"daily_connections"`
		DailyMessages    int `yaml:"daily_messages"`
	} `yaml:"limits"`

	Timing struct {
		MinDelayMs int `yaml:"min_delay_ms"`
		MaxDelayMs int `yaml:"max_delay_ms"`
	} `yaml:"timing"`

	LinkedInEmail    string
	LinkedInPassword string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),
	}

	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
