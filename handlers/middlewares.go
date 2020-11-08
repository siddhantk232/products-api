package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/siddhantk232/products-api/data"
)

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
