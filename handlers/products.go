package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siddhantk232/go-micro/data"
)

// ProductKey use it to access product from r.Context()
type ProductKey struct{}

// Products handler
type Products struct {
	log *log.Logger
}

// NewProducts create a new products handler
func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

// GetProducts GET all the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "parse error", http.StatusInternalServerError)
	}
}

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

// Middlewares

// ParseProductBody middleware parses the request body as a valid Product struct
func (p *Products) ParseProductBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := product.FromJSON(r.Body)

		if err != nil {
			p.log.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Can't create product", http.StatusBadRequest)
			return
		}

		// validate product

		err = product.Validate()

		if err != nil {
			p.log.Println("[ERROR] invalid product from body", err)
			http.Error(
				rw,
				fmt.Sprintf("Can't create product: %s", err),
				http.StatusBadRequest,
			)
			return

		}

		ctx := context.WithValue(r.Context(), ProductKey{}, product)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
