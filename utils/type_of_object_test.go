package utils_test

import (
	"testing"

	"github.com/bvestli/check_endoflife/utils"
)

func TestTypeofObject(t *testing.T) {
	test_data := []struct {
		input    interface{}
		expected string
	}{
		{true, "boolean"},
		{"string", "string"},
		{2, "unknown"},
	}

	for _, td := range test_data {
		actual := utils.TypeofObject(td.input)
		if actual != td.expected {
			t.Errorf("Expected: %s\nGot: %s", td.expected, actual)
		}
	}
}
