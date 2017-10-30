package csvx

import (
	"encoding/json"
	"errors"
)

func Top(n int, h Handler) Handler {
	if n <= 0 {
		return h
	}
	return func(line int, record interface{}) (bool, error) {
		if line >= n {
			return false, nil
		}
		return h(line, record)
	}
}

func JsonH(decoder func(data []byte) (interface{}, error), h Handler) Handler {
	if decoder == nil {
		decoder = ToFreejson
	}
	var header []string
	return func(line int, record0 interface{}) (bool, error) {
		record, ok := record0.([]string)
		if !ok {
			return false, errors.New("The first line is not header")
		}
		if len(record) == 0 {
			return false, errors.New("The record is empty")
		}
		if line == 0 {
			header = record
			return true, nil
		}
		recordLen := len(record)
		jsonRecord := make(map[string]string, len(header))
		for i, field := range header {
			var s string
			if i < recordLen {
				s = record[i]
			}
			jsonRecord[field] = s
		}
		data, err := json.Marshal(jsonRecord)
		if err != nil {
			return false, err
		}
		v, err := decoder(data)
		if err != nil {
			return false, err
		}
		return h(line, v)
	}
}
