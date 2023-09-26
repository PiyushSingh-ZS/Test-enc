package models

type Order struct {
	ID       int    `json:"ID"`
	Item     string `json:"Item"`
	Quantity int    `json:"quantity"`
}
