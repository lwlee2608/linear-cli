package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var issueCommentCmd = &cobra.Command{
	Use:   "comment <id> <body>",
	Short: "Add a comment to an issue",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		comment, err := service.AddComment(context.Background(), args[0], args[1])
		if err != nil {
			return err
		}

		fmt.Printf("Comment added to %s\n", args[0])
		if comment.User != nil {
			fmt.Printf("Author:  %s\n", comment.User.Name)
		}
		fmt.Printf("Body:    %s\n", comment.Body)
		return nil
	},
}

var issueCommentEditCmd = &cobra.Command{
	Use:   "edit <comment-id> <body>",
	Short: "Edit a comment",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		comment, err := service.EditComment(cmd.Context(), args[0], args[1])
		if err != nil {
			return err
		}

		fmt.Printf("Comment %s updated\n", comment.ID)
		if comment.User != nil {
			fmt.Printf("Author:  %s\n", comment.User.Name)
		}
		fmt.Printf("Body:    %s\n", comment.Body)
		return nil
	},
}

func init() {
	issueCommentCmd.AddCommand(issueCommentEditCmd)
}
