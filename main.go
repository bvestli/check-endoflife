package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bvestli/check_endoflife/product"
	"github.com/bvestli/check_endoflife/utils"
)

func main() {
	productNames, err := utils.GetProductsFromFile("products.yaml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var productData []product.Product
	for productName, productVersion := range productNames {
		singleProductData, err := product.FullProductData(productName, productVersion)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		productData = append(productData, singleProductData)
	}

	fmt.Println(utils.PrettyPrintJSON(productData))
}
