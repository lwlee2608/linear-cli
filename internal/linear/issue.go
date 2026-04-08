package linear

import (
	"context"

	"github.com/lwlee2608/linear-cli/pkg/linear"
)

func (s *Service) SearchIssues(ctx context.Context, query string) ([]linear.Issue, error) {
	result, err := s.client.SearchIssues(ctx, query, 20, "")
	if err != nil {
		return nil, err
	}
	return result.Nodes, nil
}
