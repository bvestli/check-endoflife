package utils

import "encoding/json"

func PrettyPrintJSON(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")

	return string(s)
}
