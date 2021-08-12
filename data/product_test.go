package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "hello",
		SKU:   "h123",
		Price: 1.1,
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
