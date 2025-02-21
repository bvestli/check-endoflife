package utils_test

import (
	"strings"
	"testing"

	"github.com/bvestli/check_endoflife/product"
	"github.com/bvestli/check_endoflife/utils"
)

func TestPrettyPrintJSON(t *testing.T) {
	product_data := []product.Product{
		{
			Name:          "product1",
			Cycle:         "2024-02-12",
			ReleaseDate:   "29.04.2025",
			EndOfLifeDate: "2.11.20",
			Latest:        "2025-04-29",
			Link:          "",
			Support:       "2.11.3",
			Discontinued:  "3.3",
			MyVersion:     "1.0",
			LastestCycle:  "2.0",
		},
		{
			Name:          "product2",
			Cycle:         "2024-03-05",
			ReleaseDate:   "30.06.2025",
			EndOfLifeDate: "10.4.15",
			Latest:        "2025-06-30",
			Link:          "",
			Support:       "10.4.14",
			Discontinued:  "11.5",
			MyVersion:     "2.0",
			LastestCycle:  "3.0",
		},
	}

	result := utils.PrettyPrintJSON(product_data)

	expected := `[
  {
    "name": "product1",
    "cycle": "2024-02-12",
    "releaseDate": "29.04.2025",
    "eol": "2.11.20",
    "latest": "2025-04-29",
    "support": "2.11.3",
    "discontinued": "3.3",
    "myversion": "1.0",
    "latestcycle": "2.0"
  },
  {
    "name": "product2",
    "cycle": "2024-03-05",
    "releaseDate": "30.06.2025",
    "eol": "10.4.15",
    "latest": "2025-06-30",
    "support": "10.4.14",
    "discontinued": "11.5",
    "myversion": "2.0",
    "latestcycle": "3.0"
  }
]`

	if result != expected {
		t.Errorf("Expected: %s\nGot: %s", strings.TrimSpace(expected), strings.TrimSpace(result))
	}

}
