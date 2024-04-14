package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}

	return string(b)
}

// Helper function to convert claim to string
func ConvertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v // Value is already a string, return it
	case float64:
		// Convert float64 to string
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		// For other types, use fmt.Sprintf to convert to string
		return fmt.Sprintf("%v", value)
	}
}
