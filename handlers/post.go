package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siddhantk232/products-api/data"
)

// AddProduct POST a product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(ProductKey{}).(*data.Product)
	data.AddProduct(product)
}

// UpdateProduct PUT the product in ProductList
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(rw, "invalid id, must be number.", http.StatusBadRequest)
	}

	product := r.Context().Value(ProductKey{}).(*data.Product)

	updateError := data.UpdateProduct(id, product)

	if updateError == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if updateError != nil {
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

}
