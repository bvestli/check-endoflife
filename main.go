package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bvestli/check_endoflife/product"
	"github.com/bvestli/check_endoflife/utils"
)

func main() {
	product_names, err := utils.GetProductsFromFile("products.yaml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var product_data []product.Product
	for product_name, product_version := range product_names {
		single_product_data, err := product.FullProductData(product_name, product_version)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		product_data = append(product_data, single_product_data)
	}

	fmt.Println(utils.PrettyPrintJSON(product_data))
}
