package hexx

import (
	"encoding/hex"
)

func EncodeString(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	return hex.EncodeToString(data)
}

func DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func DecodeStringDef(s string, def []byte) []byte {
	data, err := DecodeString(s)
	if err != nil {
		return def
	}
	return data
}
