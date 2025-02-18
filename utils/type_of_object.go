package utils

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
