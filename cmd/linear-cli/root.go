package main

import (
	"github.com/spf13/cobra"

	internallinear "github.com/lwlee2608/linear-cli/internal/linear"
	"github.com/lwlee2608/linear-cli/pkg/linear"
)

var service *internallinear.Service

var rootCmd = &cobra.Command{
	Use:   "linear-cli",
	Short: "CLI for Linear",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		InitConfig()
		client := linear.NewClient(config.Linear.APIKey)
		service = internallinear.NewService(client)
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)
}
