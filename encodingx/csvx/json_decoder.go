package csvx

import (
	"encoding/json"
	"github.com/cnych/stardust/encodingx/jsonx/freejson"
)

func ToFreejson(data []byte) (interface{}, error) {
	return freejson.Unmarshal(data, nil)
}

func ToVal(v interface{}) func(data []byte) (interface{}, error) {
	return func(data []byte) (interface{}, error) {
		err := json.Unmarshal(data, v)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
