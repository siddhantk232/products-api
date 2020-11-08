package handlers

import (
	"net/http"

	"github.com/siddhantk232/products-api/data"
)

// GetProducts GET all the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "parse error", http.StatusInternalServerError)
	}
}
