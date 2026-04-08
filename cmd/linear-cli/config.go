package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/lwlee2608/adder"
	"github.com/lwlee2608/linear-cli/pkg/linear"
)

type Config struct {
	Linear linear.Config
}

var config Config

func InitConfig() error {
	adder.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	adder.AutomaticEnv()

	if err := adder.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	configJSON, err := adder.PrettyJSON(config)
	if err == nil {
		slog.Debug("Config loaded:")
		slog.Debug(configJSON)
	}

	if config.Linear.APIKey == "" {
		return fmt.Errorf("LINEAR_API_KEY environment variable is required")
	}
	return nil
}
