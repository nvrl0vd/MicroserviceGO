package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Gera",
		Price: 12.00,
		SKU:   "asd-123",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
