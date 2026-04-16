package main

import (
	"context"
	"fmt"
	"strings"
	"time"

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
		if issue.Project != nil {
			fmt.Printf("Project:     %s\n", issue.Project.Name)
		}
		fmt.Printf("Created:     %s\n", issue.CreatedAt.Format(time.RFC3339))
		fmt.Printf("Priority:    %d\n", issue.Priority)
		if issue.Assignee != nil {
			fmt.Printf("Assignee:    %s\n", issue.Assignee.Name)
		}
		if len(issue.Labels.Nodes) > 0 {
			names := make([]string, len(issue.Labels.Nodes))
			for i, l := range issue.Labels.Nodes {
				names[i] = l.Name
			}
			fmt.Printf("Labels:      %s\n", strings.Join(names, ", "))
		}
		if issue.Description != "" {
			fmt.Printf("Description: %s\n", issue.Description)
		}
		return nil
	},
}
