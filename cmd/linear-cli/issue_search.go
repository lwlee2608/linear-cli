package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var issueSearchCmd = &cobra.Command{
	Use:   "search <keywords>",
	Short: "Search issues by keywords",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.Join(args, " ")

		issues, err := service.SearchIssues(context.Background(), query)
		if err != nil {
			return err
		}

		if len(issues) == 0 {
			fmt.Println("No issues found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTITLE\tSTATE\tASSIGNEE")
		for _, issue := range issues {
			assignee := "-"
			if issue.Assignee != nil {
				assignee = issue.Assignee.Name
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
				issue.Identifier,
				truncate(issue.Title, 60),
				issue.State.Name,
				assignee,
			)
		}
		return w.Flush()
	},
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
