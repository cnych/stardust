package unicodex

import (
	"errors"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

func EncodingOf(charset string) encoding.Encoding {
	switch strings.ToLower(charset) {
	case "gbk":
		return simplifiedchinese.GBK
	case "gb18030":
		return simplifiedchinese.GB18030
	// TODO: Other encoding
	default:
		return nil
	}
}

func isUtf8(charset string) bool {
	if charset == "" {
		return true
	}
	charset = strings.ToLower(charset)
	return charset == "utf-8" || charset == "utf8"
}

func Decode(data []byte, charset string) (string, error) {
	if data == nil {
		return "", errors.New("Nil data")
	}
	if isUtf8(charset) {
		return string(data), nil
	}
	encoding := EncodingOf(charset)
	if encoding == nil {
		return "", fmt.Errorf("Not found encoding %s", charset)
	}
	buff, err := encoding.NewDecoder().Bytes(data)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func Encode(utf8Str string, charset string) ([]byte, error) {
	if isUtf8(charset) {
		return []byte(utf8Str), nil
	}
	encoding := EncodingOf(charset)
	if encoding == nil {
		return nil, fmt.Errorf("Not found encoding %s", charset)
	}
	return encoding.NewEncoder().Bytes([]byte(utf8Str))
}
