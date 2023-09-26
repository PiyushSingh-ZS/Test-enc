package inventory

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Get(ctx *gofr.Context, item string) (int, error) {
	count := len(item)
	return count, nil
}
