package stringsx

import "strings"

// returns the value of the map for a certain key.
// Ignore case when comparing if case insensitive.
func GetMapValue(m map[string]string, key string, caseInsensitive bool) string {
	if caseInsensitive {
		for k, v := range m {
			if strings.ToLower(k) == strings.ToLower(key) {
				return v
			}
		}
	}
	return m[key]
}

// transforms each item of a map to lowercase.
func MapKeyToLower(m map[string]string) {
	temp := make(map[string]string, len(m))
	for key, value := range m {
		temp[strings.ToLower(key)] = value
		delete(m, key)
	}
	for key, value := range temp {
		m[key] = value
	}
}
