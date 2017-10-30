package typemarshal

import (
	"encoding/json"
	"errors"
	"reflect"
)

type typeValue struct {
	T string          `json:"t"`
	V json.RawMessage `json:"v"`
}

type Marshaler func(v interface{}) ([]byte, error)
type Unmarshaler func(data []byte) (interface{}, error)

type encoder struct {
	Marshal   Marshaler
	Unmarshal Unmarshaler
}

var (
	typeToName = map[reflect.Type]string{nil: "nil"}
	nameToType = map[string]reflect.Type{}
	encoders   = map[string]encoder{"nil": encoder{nilMarshaler, nilUnmarshaler}}
)

func Register(name string, typ reflect.Type) error {
	return RegisterCustom(name, typ, DefaultMarshaler, TypeUnmarshalerOf(typ))
}

func RegisterCustom(name string, typ reflect.Type, marshal Marshaler, unmarshal Unmarshaler) error {
	if name == "" {
		return errors.New("Missing name")
	}
	if typ == nil {
		name = "nil"
	}
	typeToName[typ] = name
	nameToType[name] = typ
	encoders[name] = encoder{marshal, unmarshal}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	typ := reflect.TypeOf(v)
	name, ok := typeToName[typ]
	if !ok {
		return nil, errors.New("Not found type name")
	}
	enc, ok := encoders[name]
	if !ok {
		return nil, errors.New("Not found marshaler")
	}
	data, err := enc.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.Marshal(&typeValue{name, json.RawMessage(data)})
}

func Unmarshal(data []byte) (interface{}, error) {
	var tv typeValue
	err := json.Unmarshal(data, &tv)
	if err != nil {
		return nil, err
	}
	enc, ok := encoders[tv.T]
	if !ok {
		return nil, errors.New("Not found marshaler")
	}
	return enc.Unmarshal([]byte(tv.V))
}

func DefaultMarshaler(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func TypeUnmarshalerOf(typ reflect.Type) Unmarshaler {
	return func(data []byte) (interface{}, error) {
		v := reflect.New(typ)
		err := json.Unmarshal(data, v.Interface())
		if err != nil {
			return nil, err
		}
		return v.Elem().Interface(), nil
	}
}

func nilMarshaler(v interface{}) ([]byte, error) {
	return []byte("null"), nil
}

func nilUnmarshaler(data []byte) (interface{}, error) {
	if string(data) != "null" {
		return nil, errors.New("Nil unmarshaler error")
	}
	return nil, nil
}
