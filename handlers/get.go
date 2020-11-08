package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siddhantk232/products-api/data"
)

// GetProduct use this to get a single product
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(rw, "invalid id, must be number.", http.StatusBadRequest)
	}

	product, err := data.GetProductByID(id)

	if err != nil {
		http.Error(rw, fmt.Sprintf("can't get product %s", err), http.StatusBadRequest)
	}

	err = data.ToJSON(product, rw)

	if err != nil {
		http.Error(rw, fmt.Sprintf("can't get product %s", err), http.StatusBadRequest)
	}

}
