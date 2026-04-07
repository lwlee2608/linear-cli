package linear

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	APIKey string `yaml:"api_key"`
}

func LoadAPIKey() (string, error) {
	if key := os.Getenv("LINEAR_API_KEY"); key != "" {
		return key, nil
	}

	cfg, err := loadConfigFile()
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no API key found: set LINEAR_API_KEY or add api_key to %s", ConfigPath())
		}
		return "", fmt.Errorf("failed to read config file %s: %w", ConfigPath(), err)
	}
	if cfg.APIKey == "" {
		return "", fmt.Errorf("no API key found: set LINEAR_API_KEY or add api_key to %s", ConfigPath())
	}
	return cfg.APIKey, nil
}

func loadConfigFile() (*config, error) {
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		return nil, err
	}
	var cfg config
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
