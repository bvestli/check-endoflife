package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Product struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Config struct {
	Products []Product `yaml:"products"`
}

func GetProductsFromFile(fileName string) (map[string]string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal yaml: %v", err)
	}

	result := make(map[string]string)
	for _, product := range config.Products {
		result[product.Name] = product.Version
	}

	return result, nil
}
