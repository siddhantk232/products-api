package setup

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// /     GET    - get all products.
// /{id} GET    - get one product.
// /     POST   - create product.
// /{id} UPDATE - update a product.
// /{id} DELETE - delete a product.

// SetupRoutes register for products-api
func SetupRoutes(sm *mux.Router) {
	sm.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(rw, "GET products")

	}).Methods("GET")

	sm.HandleFunc("/{id:[0-9]+}", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(rw, "GET one product")

	}).Methods("GET")

	sm.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(rw, "create product")

	}).Methods("POST")

	sm.HandleFunc("/{id:[0-9]+}", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(rw, "UPDATE product")

	}).Methods("UPDATE")

	sm.HandleFunc("/{id:[0-9]+}", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(rw, "DELETE product")

	}).Methods("DELETE")
}
