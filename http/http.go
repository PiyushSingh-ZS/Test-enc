package http

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"inventory/services"
)

type Handler struct {
	service services.OrderSvc
}

func New(service services.OrderSvc) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Index(ctx *gofr.Context) (interface{}, error) {
	item := ctx.Param("item")
	return h.service.Get(ctx, item)
}
