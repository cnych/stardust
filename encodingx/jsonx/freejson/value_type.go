package freejson

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type ValueType int

const (
	UnknownField ValueType = 0
	NoField      ValueType = 1
	NullField    ValueType = 2
	NumField     ValueType = 3
	BoolField    ValueType = 4
	StrField     ValueType = 5
	ArrayField   ValueType = 6
	ObjectField  ValueType = 7
)

func TypeOf(v interface{}) ValueType {
	if v == nil {
		return NullField
	}
	switch v.(type) {
	case string:
		return StrField
	case int, int32, int64, float64, float32, json.Number:
		return NumField
	case bool:
		return BoolField
	case Array, []interface{}:
		return ArrayField
	case Object, map[string]interface{}, bson.M, bson.D:
		return ObjectField
	default:
		return UnknownField
	}
}
