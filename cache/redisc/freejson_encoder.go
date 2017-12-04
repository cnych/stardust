package redisc

import "github.com/cnych/starjazz/encodingx/jsonx/freejson"

var (
	FreejsonEncoder     = &freejsonEncoder{}
	GzipFreejsonEncoder = &GzipEncoder{Encoder: FreejsonEncoder}
)

type freejsonEncoder struct {
}

func (enc *freejsonEncoder) Encode(v interface{}) ([]byte, error) {
	return freejson.Marshal(v)
}

func (enc *freejsonEncoder) Decode(data []byte) (interface{}, error) {
	return freejson.Unmarshal(data, nil)
}
