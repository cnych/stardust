package freejson

import (
	"github.com/cnych/stardust/encodingx/jsonx"
	"time"
)

type Array []interface{}

func (a Array) Len() int {
	return len(a)
}

func (a Array) Has(index int) bool {
	return index >= 0 && index < len(a)
}

func (a Array) Each(f func(index int, v interface{})) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Array) EachIndex(f func(index int)) {
	for i, _ := range a {
		f(i)
	}
}

func (a Array) EachElem(f func(v interface{})) {
	for _, v := range a {
		f(v)
	}
}

func (a Array) Filter(result Array, pred func(v interface{}) bool) Array {
	if result == nil {
		result = make(Array, 0, 4)
	}
	for _, v := range a {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func (a Array) FieldType(index int) ValueType {
	if a == nil {
		return NoField
	}
	if !a.Has(index) {
		return NoField
	}
	return TypeOf(a[index])
}

func (a Array) IntfAt(index int, def interface{}) interface{} {
	if !a.Has(index) {
		return def
	}
	return a[index]
}

func (a Array) StrAt(index int, def string) string {
	return ToStr(a.IntfAt(index, nil), def)
}

func (a Array) AsStrAt(index int, def string) string {
	return AsStr(a.IntfAt(index, nil), def)
}

func (a Array) BoolAt(index int, def bool) bool {
	return ToBool(a.IntfAt(index, nil), def)
}

func (a Array) IntAt(index, def int) int {
	return ToInt(a.IntfAt(index, nil), def)
}

func (a Array) Int64At(index int, def int64) int64 {
	return ToInt64(a.IntfAt(index, nil), def)
}

func (a Array) Float64At(index int, def float64) float64 {
	return ToFloat64(a.IntfAt(index, nil), def)
}

func (a Array) ObjectAt(index int, def Object) Object {
	return ToObject(a.IntfAt(index, nil), def)
}

func (a Array) ArrayAt(index int, def Array) Array {
	return ToArray(a.IntfAt(index, nil), def)
}

func (a Array) AsArrayAt(index int, def Array) Array {
	return AsArray(a.IntfAt(index, nil), def)
}

func (a Array) AsStringArrayAt(index int, def []string) []string {
	return AsStringArray(a.IntfAt(index, nil), def)
}

func (a Array) TimeAt(index int, def time.Time) time.Time {
	return ToTime(a.IntfAt(index, nil), def)
}

func (a Array) AsTimeAt(index int, def time.Time) time.Time {
	return AsTime(a.IntfAt(index, nil), def)
}

func (a Array) Set(index int, v interface{}) {
	if a == nil {
		return
	}
	a[index] = v
}

func (a Array) Clear() Array {
	if a == nil {
		return nil
	}
	return Array(a[:0])
}

func (a Array) Add(vs ...interface{}) Array {
	return Array(append(a, vs...))
}

func (a Array) RemoveIf(pred func(index int, v interface{}) bool) Array {
	if a == nil {
		return nil
	}
	a1 := make(Array, 0, len(a))
	for i, v := range a {
		if pred(i, v) {
			a1 = append(a1, v)
		}
	}
	return a1
}

func (a Array) String() string {
	return jsonx.MarshalIndentString(a, "")
}
