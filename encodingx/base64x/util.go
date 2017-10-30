package base64x

import (
	"encoding/base64"
)

func EncodeString(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

func DecodeString(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

func DecodeStdStr(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
