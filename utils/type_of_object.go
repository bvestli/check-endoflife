package utils

// TypeofObject takes an interface and returns the type of the object as a string
func TypeofObject(variable interface{}) string {
	switch variable.(type) {
	case bool:
		return "boolean"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
