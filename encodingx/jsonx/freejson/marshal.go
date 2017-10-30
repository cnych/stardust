package freejson

import (
	"bytes"
	"encoding/json"
)

type ValueDecoder func(interface{}) (interface{}, error)

func Unmarshal(data []byte, valDec ValueDecoder) (interface{}, error) {
	if valDec == nil {
		valDec = func(v interface{}) (interface{}, error) {
			return v, nil
		}
	}

	var v interface{}
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.UseNumber()
	err := dec.Decode(&v)
	if err != nil {
		return nil, err
	}
	var trans func(v interface{}) (interface{}, error)
	trans = func(v interface{}) (interface{}, error) {
		switch v1 := v.(type) {
		case map[string]interface{}:
			for k, e := range v1 {
				e1, err := trans(e)
				if err != nil {
					return nil, err
				}
				v1[k] = e1
			}
			return valDec(Object(v1))
		case []interface{}:
			for i, e := range v1 {
				e1, err := trans(e)
				if err != nil {
					return nil, err
				}
				v1[i] = e1
			}
			return valDec(Array(v1))
		default:
			return valDec(v)
		}
	}
	return trans(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func MarshalString(v interface{}, def string) string {
	data, err := Marshal(v)
	if err != nil {
		return def
	}
	return string(data)
}

func MarshalIndentString(v interface{}, def string) string {
	data, err := MarshalIndent(v, "", "  ")
	if err != nil {
		return def
	}
	return string(data)
}

func FromFreejson(freejsonData interface{}, v interface{}) error {
	data, err := json.Marshal(freejsonData)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func ToFreejson(v interface{}) (interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return Unmarshal(data, nil)
}
