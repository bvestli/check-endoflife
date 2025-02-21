package utils

import "encoding/json"

// PrettyPrintJSON takes an interface and returns a pretty printed JSON string
func PrettyPrintJSON(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")

	return string(s)
}
