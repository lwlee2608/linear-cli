package main

import (
	"github.com/spf13/cobra"

	internallinear "github.com/lwlee2608/linear-cli/internal/linear"
	"github.com/lwlee2608/linear-cli/pkg/linear"
)

var service *internallinear.Service

func initService(cmd *cobra.Command, args []string) error {
	if err := InitConfig(); err != nil {
		return err
	}
	client := linear.NewClient(config.Linear.APIKey)
	service = internallinear.NewService(client)
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "linear",
	Short: "CLI for Linear",
}

func init() {
	rootCmd.AddCommand(issueCmd)
}
