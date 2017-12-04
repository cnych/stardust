package redisc

import (
	"github.com/cnych/starjazz/encodingx/jsonx/typemarshal"
)

var (
	TypeMarshalEncoder     = &typeMarshalEncoder{}
	GzipTypeMarshalEncoder = &GzipEncoder{TypeMarshalEncoder}
)

type typeMarshalEncoder struct {
}

func (enc *typeMarshalEncoder) Encode(v interface{}) ([]byte, error) {
	return typemarshal.Marshal(v)
}

func (enc *typeMarshalEncoder) Decode(data []byte) (interface{}, error) {
	return typemarshal.Unmarshal(data)
}
