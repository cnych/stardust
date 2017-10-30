package jsonx

import "encoding/json"

func MarshalString(v interface{}, def string) string {
	data, err := json.Marshal(v)
	if err != nil {
		return def
	}
	return string(data)
}

func MarshalIndentString(v interface{}, def string) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return def
	}
	return string(data)
}

func Format(data []byte) ([]byte, error) {
	var o interface{}
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(o, "", "  ")
}

func FormatDef(data, def string) string {
	data1, err := Format([]byte(data))
	if err != nil {
		return def
	}
	return string(data1)
}

func IsValid(data []byte) bool {
	var v interface{}
	err := json.Unmarshal(data, &v)
	return err != nil
}
