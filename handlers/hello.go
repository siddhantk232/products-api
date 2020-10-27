package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello routes handler
type Hello struct {
	log *log.Logger
}

// NewHello use this to create the Hello handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// GetHello greet users to the api
func (h *Hello) GetHello(rw http.ResponseWriter, r *http.Request) {

	data, error := ioutil.ReadAll(r.Body)

	if error != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
