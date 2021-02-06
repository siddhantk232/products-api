package products

import (
	"fmt"
	"log"
	"net/http"
)

type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger: l}
}

func (p *ProductsHandler) ListProducts(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Listing products...")
}

func (p *ProductsHandler) ListProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Listing a product...")
}

func (p *ProductsHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "creating a product...")
}

func (p *ProductsHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "updating a product...")
}

func (p *ProductsHandler) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "deleting a product...")
}
