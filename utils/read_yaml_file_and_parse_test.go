package utils_test

import (
	"testing"

	"github.com/bvestli/check_endoflife/utils"
)

func TestGetProductsFromFile(t *testing.T) {
	result := make(map[string]string)

	result, _ = utils.GetProductsFromFile("../testdata/products.yaml")

	expected := map[string]string{
		"product1": "1.0",
		"product2": "2.0",
		"product3": "3.0",
	}

	for key, value := range expected {
		if result[key] != value {
			t.Errorf("Expected: %s\nGot: %s", value, result[key])
		}
	}
}
