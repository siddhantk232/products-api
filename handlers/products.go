package handlers

import (
	"log"

	"github.com/siddhantk232/currency/protos/currency"
)

// ProductKey use it to access product from r.Context()
type ProductKey struct{}

// Products handler
type Products struct {
	log *log.Logger
	cc  currency.CurrencyClient
}

// NewProducts create a new products handler
func NewProducts(log *log.Logger, cc currency.CurrencyClient) *Products {
	return &Products{log, cc}
}
