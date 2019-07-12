package freejson

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cnych/stardust/timex"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func ToStr(v interface{}, def string) string {
	if v == nil {
		return def
	}
	return v.(string)
}

func AsStr(v interface{}, def string) string {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case string:
		return v1
	case bson.ObjectId:
		return v1.Hex()
	case int, int64, bool, float64, float32:
		return fmt.Sprintf("%v", v1)
	case time.Time:
		return v1.Format(time.RFC3339)
	default:
		return MarshalString(v, def)
	}
}

func ToBool(v interface{}, def bool) bool {
	if v == nil {
		return def
	}
	return v.(bool)
}

func AsBool(v interface{}, def bool) bool {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case string:
		return v1 == "true" || v1 != "0"
	case int, int64, float64, float32:
		return v1 == 1
	case bool:
		return v1
	default:
		return false
	}
}

func BoolAsInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func ToInt(v interface{}, def int) int {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case int:
		return v1
	case int64:
		return int(v1)
	case int32:
		return int(v1)
	case float64:
		return int(v1)
	case float32:
		return int(v1)
	case json.Number:
		i64, err := v1.Int64()
		if err != nil {
			panic(err)
		}
		return int(i64)
	default:
		panic(errors.New("Not as int"))
		return 0
	}
}

func ToInt64(v interface{}, def int64) int64 {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case int:
		return int64(v1)
	case int64:
		return v1
	case int32:
		return int64(v1)
	case float64:
		return int64(v1)
	case float32:
		return int64(v1)
	case json.Number:
		i64, err := v1.Int64()
		if err != nil {
			panic(err)
		}
		return i64
	default:
		panic(errors.New("Not as int64"))
		return 0
	}
}

func ToFloat64(v interface{}, def float64) float64 {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case int:
		return float64(v1)
	case int64:
		return float64(v1)
	case int32:
		return float64(v1)
	case float64:
		return v1
	case float32:
		return float64(v1)
	case json.Number:
		f64, err := v1.Float64()
		if err != nil {
			panic(err)
		}
		return f64
	default:
		panic(errors.New("Not as float64"))
		return 0
	}
}

func ToArray(v interface{}, def Array) Array {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case []interface{}:
		return Array(v1)
	default:
		return v.(Array)
	}
}

func ToObject(v interface{}, def Object) Object {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case map[string]interface{}:
		return Object(v1)
	case bson.M:
		return Object(v1)
	case bson.D:
		return Object(v1.Map())
	default:
		return v.(Object)
	}
}

func AsArray(v interface{}, def Array) Array {
	if v == nil {
		return def
	}
	switch a1 := v.(type) {
	case Array:
		return a1
	case []interface{}:
		return Array(a1)
	default:
		return Array{v}
	}
}

func AsStringArray(v interface{}, def []string) []string {
	arr := AsArray(v, nil)
	if v == nil {
		return def
	}
	strArr := make([]string, 0, len(arr))
	for _, v := range arr {
		strArr = append(strArr, AsStr(v, ""))
	}
	return strArr
}

func ToTime(v interface{}, def time.Time) time.Time {
	if v == nil {
		return def
	}
	return v.(time.Time)
}

func AsTime(v interface{}, def time.Time) time.Time {
	if v == nil {
		return def
	}
	switch v1 := v.(type) {
	case time.Time:
		return v1
	case string:
		t, err := timex.TryParse(v1)
		if err != nil {
			return def
		}
		return t
	case int32:
		return timex.FromEpoch(int64(v1))
	case int64:
		return timex.FromEpoch(v1)
	case json.Number:
		epoch, err := v1.Int64()
		if err != nil {
			return def
		}
		return timex.FromEpoch(epoch)
	default:
		return def
	}
}
