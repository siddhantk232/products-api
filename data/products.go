package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/siddhantk232/currency/protos/currency"
)

// Product struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// ProductsDB basic product interaction
type ProductsDB struct {
	log log.Logger
	cc  currency.CurrencyClient
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
func (p *ProductsDB) GetProducts(destCurrency string) (Products, error) {

	if destCurrency == "" {
		return ProductList, nil
	}

	rate, err := p.getRate(destCurrency)

	if err != nil {
		p.log.Println("Error getting exchange rate", err)
		return nil, err
	}

	products := Products{}

	for _, product := range ProductList {
		np := *product
		np.Price = np.Price * rate

		products = append(products, &np)
	}

	return products, nil

}

// GetProductByID return a product
func (p *ProductsDB) GetProductByID(id int, destCurrency string) (*Product, error) {
	product, _, err := findProduct(id)

	if id == -1 {
		return nil, ErrorProductNotFound
	}

	if destCurrency == "" {
		return product, err

	}

	rate, err := p.getRate(destCurrency)

	if err != nil {
		p.log.Println("Error getting exchange rate", err)
		return nil, err
	}

	np := *product

	np.Price = np.Price * rate

	return &np, nil
}

// AddProduct adds a product to the ProductList
func (p *ProductsDB) AddProduct(product *Product) {
	product.ID = getNextID()
	ProductList = append(ProductList, product)
}

// UpdateProduct updates the product in the ProductList
func (p *ProductsDB) UpdateProduct(id int, product *Product) error {
	prod, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = prod.ID
	ProductList[pos] = product
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

func (p *ProductsDB) getRate(destCurrency string) (float64, error) {
	rr := &currency.RateRequest{
		Base:        currency.Currencies(currency.Currencies_value["EUR"]),
		Destination: currency.Currencies(currency.Currencies_value[destCurrency]),
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	return resp.Rate, err
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
