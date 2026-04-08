package linear

import (
	"github.com/lwlee2608/linear-cli/pkg/linear"
)

type Service struct {
	client *linear.Client
}

func NewService(client *linear.Client) *Service {
	return &Service{client: client}
}
