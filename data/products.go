package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float32
	SKU         string //A code that is assigned to each product within the company to facilitate accounting.
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		Id:          1,
		Name:        "Shirt",
		Description: "White oversized Shirt",
		Price:       799,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		Id:          1,
		Name:        "Jeans",
		Description: "Black wide-leg jeans",
		Price:       799,
		SKU:         "cde456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
