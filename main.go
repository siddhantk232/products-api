package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, error := ioutil.ReadAll(r.Body)

		if error != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.ListenAndServe(":9090", nil)
}
