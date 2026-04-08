package linear

import (
	"context"

	"github.com/lwlee2608/linear-cli/pkg/linear"
)

func (s *Service) GetIssue(ctx context.Context, id string) (*linear.Issue, error) {
	return s.client.GetIssue(ctx, id)
}

func (s *Service) SearchIssues(ctx context.Context, query string, limit int) ([]linear.Issue, error) {
	result, err := s.client.SearchIssues(ctx, query, limit, "")
	if err != nil {
		return nil, err
	}
	return result.Nodes, nil
}
