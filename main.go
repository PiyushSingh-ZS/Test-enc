package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	orderHandler "inventory/http"
	orderService "inventory/services/inventory"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	orderSvc := orderService.New()
	orderHTTP := orderHandler.New(orderSvc)

	app.REST("inventory", orderHTTP)
	app.Start()
}
