package product_test

import (
	"testing"

	"github.com/bvestli/check_endoflife/product"
)

// TestGetSingleCycle tests the GetSingleCycle function
func TestGetSingleCycle(t *testing.T) {
	singleCycleTests, err := product.GetSingleCycle("prometheus", "3.0")
	if err != nil {
		t.Errorf("GetSingleCycle(prometheus, 3.0) failed: %v", err)
	}

	if singleCycleTests.Name != "prometheus" {
		t.Errorf("GetSingleCycle(prometheus, 3.0) failed: expected Name to be prometheus, got %s", singleCycleTests.Name)
	}

	if singleCycleTests.Latest != "3.0.1" {
		t.Errorf("GetSingleCycle(prometheus, 3.0) failed: expected Latest to be 3.0.1, got %s", singleCycleTests.Cycle)
	}

	if singleCycleTests.ReleaseDate != "2024-11-14" {
		t.Errorf("GetSingleCycle(prometheus, 3.0) failed: expected ReleaseDate to be 2024-11-14, got %s", singleCycleTests.ReleaseDate)
	}
}

func TestGetLatestCycle(t *testing.T) {
	_, err := product.GetLatestCycle("prometheus")
	if err != nil {
		t.Errorf("GetLatestCycle(prometheus) failed: %v", err)
	}
}
