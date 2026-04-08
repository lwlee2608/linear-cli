package main

import "github.com/spf13/cobra"

var issueCmd = &cobra.Command{
	Use:               "issue",
	Short:             "Manage issues",
	PersistentPreRunE: initService,
}

func init() {
	issueCmd.AddCommand(issueGetCmd)
	issueCmd.AddCommand(issueSearchCmd)
}
