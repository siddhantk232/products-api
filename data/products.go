package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON decodes json
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// Validate use this to validate instance of this struct
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
		matches := re.FindAllString(fl.Field().String(), -1)

		if len(matches) != 1 {
			return false
		}
		return true
	})
	return validate.Struct(p)
}

// Products type
type Products []*Product

// ToJSON converts ProductList to json
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// GetProducts returns product list
func GetProducts() Products {
	return ProductList
}

// AddProduct adds a product to the ProductList
func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

// UpdateProduct updates the product in the ProductList
func UpdateProduct(id int, p *Product) error {
	prod, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = prod.ID
	ProductList[pos] = p
	return nil
}

// ErrorProductNotFound formatted error
var ErrorProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

func getNextID() int {
	return ProductList[len(ProductList)-1].ID + 1
}

// ProductList static list
var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.44,
		SKU:         "coffee-101",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "strong coffee",
		Price:       1.23,
		SKU:         "coffee-102",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
