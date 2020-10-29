package data

import "testing"

func TestChecksValidation(t *testing.T) {

	product := &Product{Name: "caffe", Price: 21.2, SKU: "coff-nam-ne"}

	if error := product.Validate(); error != nil {
		t.Fatal(error)
	}

}
