package setup

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(sm *mux.Router) {
	r := sm.PathPrefix("/products/").Subrouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(rw, "Hello World!")

	}).Methods("GET")

}
