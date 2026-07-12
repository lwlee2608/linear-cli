package linear

import (
	"context"

	"github.com/lwlee2608/linear-cli/pkg/linear"
)

func (s *Service) GetIssue(ctx context.Context, id string) (*linear.Issue, error) {
	return s.client.GetIssue(ctx, id)
}

func (s *Service) AddComment(ctx context.Context, issueID string, body string) (*linear.Comment, error) {
	issue, err := s.client.GetIssue(ctx, issueID)
	if err != nil {
		return nil, err
	}
	return s.client.CreateComment(ctx, linear.CommentCreateInput{IssueID: issue.ID, Body: body})
}

func (s *Service) EditComment(ctx context.Context, commentID string, body string) (*linear.Comment, error) {
	return s.client.UpdateComment(ctx, commentID, linear.CommentUpdateInput{Body: body})
}

func (s *Service) SearchIssues(ctx context.Context, query string, limit int) ([]linear.Issue, error) {
	result, err := s.client.SearchIssues(ctx, query, limit, "")
	if err != nil {
		return nil, err
	}
	return result.Nodes, nil
}
