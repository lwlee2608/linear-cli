package linear

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIKey string `yaml:"api_key"`
}

func LoadAPIKey() (string, error) {
	if key := os.Getenv("LINEAR_API_KEY"); key != "" {
		return key, nil
	}

	cfg, err := loadConfigFile()
	if err != nil {
		return "", fmt.Errorf("no API key found: set LINEAR_API_KEY or add api_key to %s", ConfigPath())
	}
	if cfg.APIKey == "" {
		return "", fmt.Errorf("no API key found: set LINEAR_API_KEY or add api_key to %s", ConfigPath())
	}
	return cfg.APIKey, nil
}

func loadConfigFile() (*Config, error) {
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ConfigPath() string {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, _ := os.UserHomeDir()
		dir = filepath.Join(home, ".config")
	}
	return filepath.Join(dir, "linear-cli", "config.yaml")
}
