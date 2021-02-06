package setup

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/siddhantk232/products-api/products"
)

// /     GET    - get all products.
// /{id} GET    - get one product.
// /     POST   - create product.
// /{id} UPDATE - update a product.
// /{id} DELETE - delete a product.

// SetupRoutes register for products-api
func SetupRoutes(sm *mux.Router, l *log.Logger) {

	productsHandler := products.NewProductsHandler(l)

	sm.HandleFunc("/", productsHandler.ListProducts).Methods("GET")

	sm.HandleFunc("/{id:[0-9]+}", productsHandler.ListProduct).Methods("GET")

	sm.HandleFunc("/", productsHandler.CreateProduct).Methods("POST")

	sm.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct).Methods("UPDATE")

	sm.HandleFunc("/{id:[0-9]+}", productsHandler.DeleteProduct).Methods("DELETE")
}
