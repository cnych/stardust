package stringsx

import (
	"github.com/cnych/stardust/netx/httpx"
	"path"
	"regexp"
	"strings"
)

func GetSuffix(s string) string {
	return strings.Replace(path.Ext(httpx.RemoveUrlScheme(s)), ".", "", 1)
}

func GetNotSpace(s ...string) string {
	for _, value := range s {
		if value != "" {
			return value
		}
	}
	return ""
}

func RemovePrefix(s, prefix string) string {
	if prefix == "" {
		return s
	}
	if !strings.HasPrefix(s, prefix) {
		return s
	}
	return s[len(prefix):]
}

func RemoveSuffix(s, suffix string) string {
	if suffix == "" {
		return s
	}
	if !strings.HasSuffix(s, suffix) {
		return s
	}
	return s[:len(s)-len(suffix)]
}

func RemoveSymbol(s, symbol string) string {
	if symbol == "" || !strings.Contains(s, symbol) {
		return s
	}
	return strings.Replace(s, symbol, "", -1)
}

func CheckPhone(s string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|6|7|8][0-9]\d{4,8})$`, s); !m {
		return false
	}
	return true
}
