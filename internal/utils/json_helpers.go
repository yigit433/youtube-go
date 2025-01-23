package utils

// getStringValue safely retrieves a string value from a nested structure
func getStringValue(data interface{}, path ...string) string {
	current := data
	for _, p := range path {
		if m, ok := current.(map[string]interface{}); ok {
			if v, exists := m[p]; exists {
				current = v
			} else {
				return ""
			}
		} else {
			return ""
		}
	}
	if str, ok := current.(string); ok {
		return str
	}
	return ""
}

// getArrayValue safely retrieves a slice of interface{} from a nested structure
func getArrayValue(data interface{}, path ...string) []interface{} {
	current := data
	for _, p := range path {
		if m, ok := current.(map[string]interface{}); ok {
			if v, exists := m[p]; exists {
				current = v
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	if arr, ok := current.([]interface{}); ok {
		return arr
	}
	return nil
}

// getFloat64Value safely retrieves a float64 value from a nested structure
func getFloat64Value(data interface{}, path ...string) float64 {
	current := data
	for _, p := range path {
		if m, ok := current.(map[string]interface{}); ok {
			if v, exists := m[p]; exists {
				current = v
			} else {
				return 0
			}
		} else {
			return 0
		}
	}
	if f, ok := current.(float64); ok {
		return f
	}
	return 0
}
