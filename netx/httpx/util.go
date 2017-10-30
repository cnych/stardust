package httpx

import "strings"

func IsURL(s string) bool {
	if s == "" {
		return false
	}
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

func EnsureURL(urlOrPath string) string {
	if strings.HasPrefix(urlOrPath, "http://") || strings.HasPrefix(urlOrPath, "https://") {
		return urlOrPath
	}
	return "http://" + urlOrPath
}

func RemoveUrlScheme(fullUrl string) string {
	if strings.HasPrefix(fullUrl, "http://") {
		return strings.TrimPrefix(fullUrl, "http://")
	}
	if strings.HasPrefix(fullUrl, "https://") {
		return strings.TrimPrefix(fullUrl, "https://")
	}
	return fullUrl
}
