package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type OrderSvc interface {
	Get(ctx *gofr.Context, item string) (int, error)
}
