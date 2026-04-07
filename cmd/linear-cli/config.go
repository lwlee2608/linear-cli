package main

import (
	"log/slog"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lwlee2608/adder"
	"github.com/lwlee2608/linear-cli/pkg/linear"
)

type Config struct {
	Linear linear.Config
}

var config Config

func InitConfig() {
	_ = godotenv.Load()

	adder.SetConfigName("application")
	adder.AddConfigPath(".")
	adder.SetConfigType("yaml")
	adder.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	adder.AutomaticEnv()

	if err := adder.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := adder.Unmarshal(&config); err != nil {
		panic(err)
	}

	configJSON, err := adder.PrettyJSON(config)
	if err == nil {
		slog.Debug("Config loaded:")
		slog.Debug(configJSON)
	}

	if config.Linear.APIKey == "" {
		panic("linear.api_key is required: set LINEAR_API_KEY env var or add it to application.yml")
	}
}
