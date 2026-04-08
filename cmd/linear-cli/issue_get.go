package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var issueGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an issue by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		issue, err := service.GetIssue(context.Background(), args[0])
		if err != nil {
			return err
		}

		fmt.Printf("ID:          %s\n", issue.Identifier)
		fmt.Printf("Title:       %s\n", issue.Title)
		fmt.Printf("State:       %s\n", issue.State.Name)
		fmt.Printf("Team:        %s\n", issue.Team.Name)
		fmt.Printf("Priority:    %d\n", issue.Priority)
		if issue.Assignee != nil {
			fmt.Printf("Assignee:    %s\n", issue.Assignee.Name)
		}
		if issue.Description != "" {
			fmt.Printf("Description: %s\n", issue.Description)
		}
		return nil
	},
}
