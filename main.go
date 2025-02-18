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

	for product_name, product_version := range product_names {
		json_data, err := product.PrintProductJSON(product_name, product_version)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println(json_data)
	}
}
