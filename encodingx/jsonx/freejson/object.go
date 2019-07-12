package freejson

import (
	"github.com/cnych/stardust/encodingx/jsonx"
	"time"
)

type Object map[string]interface{}

func (o Object) Len() int {
	return len(o)
}

func (o Object) Has(f string) bool {
	if o == nil {
		return false
	}
	_, ok := o[f]
	return ok
}

func (o Object) Keys() []string {
	keys := make([]string, 0, len(o))
	for k, _ := range o {
		keys = append(keys, k)
	}
	return keys
}

func (o Object) Each(f func(k string, v interface{})) {
	for k, v := range o {
		f(k, v)
	}
}

func (o Object) EachKey(f func(k string)) {
	for k, _ := range o {
		f(k)
	}
}

func (o Object) EachValue(f func(v interface{})) {
	for _, v := range o {
		f(v)
	}
}

func (o Object) Filter(result Object, pred func(k string, v interface{}) bool) Object {
	if result == nil {
		result = Object{}
	}
	for k, v := range o {
		if pred(k, v) {
			result[k] = v
		}
	}
	return result
}

func (o Object) FieldType(f string) ValueType {
	if o == nil {
		return NoField
	}
	v, ok := o[f]
	if !ok {
		return NoField
	}
	return TypeOf(v)
}

func (o Object) IntfField(f string, def interface{}) interface{} {
	if o == nil {
		return def
	}
	v, ok := o[f]
	if !ok {
		return def
	}
	return v
}

func (o Object) StrField(f, def string) string {
	return ToStr(o.IntfField(f, nil), def)
}

func (o Object) AsStrField(f, def string) string {
	return AsStr(o.IntfField(f, nil), def)
}

func (o Object) BoolField(f string, def bool) bool {
	return ToBool(o.IntfField(f, nil), def)
}

func (o Object) AsBoolField(f string, def bool) bool {
	return AsBool(o.IntfField(f, nil), def)
}

func (o Object) BoolFieldAsInt(f string, def bool) int {
	return BoolAsInt(o.BoolField(f, def))
}

func (o Object) IntField(f string, def int) int {
	return ToInt(o.IntfField(f, nil), def)
}

func (o Object) Int64Field(f string, def int64) int64 {
	return ToInt64(o.IntfField(f, nil), def)
}

func (o Object) Float64Field(f string, def float64) float64 {
	return ToFloat64(o.IntfField(f, nil), def)
}

func (o Object) ObjectField(f string, def Object) Object {
	return ToObject(o.IntfField(f, nil), def)
}

func (o Object) ArrayField(f string, def Array) Array {
	return ToArray(o.IntfField(f, nil), def)
}

func (o Object) AsArrayField(f string, def Array) Array {
	return AsArray(o.IntfField(f, nil), def)
}

func (o Object) AsStringArrayField(f string, def []string) []string {
	return AsStringArray(o.IntfField(f, nil), def)
}

func (o Object) TimeField(f string, def time.Time) time.Time {
	return ToTime(o.IntfField(f, nil), def)
}

func (o Object) AsTimeField(f string, def time.Time) time.Time {
	return AsTime(o.IntfField(f, nil), def)
}

func (o Object) Set(f string, v interface{}) {
	if o == nil {
		return
	}
	o[f] = v
}

func (o Object) Ensure(f string, v interface{}) {
	if o == nil {
		return
	}
	_, ok := o[f]
	if ok {
		return
	}
	o[f] = v
}

func (o Object) EnsureF(f string, vf func(f string) (interface{}, error)) error {
	if o == nil {
		return nil
	}
	_, ok := o[f]
	if ok {
		return nil
	}
	v, err := vf(f)
	if err != nil {
		return err
	}
	o[f] = v
	return nil
}

func (o Object) SetNotNil(f string, v interface{}) {
	if o == nil {
		return
	}
	if v != nil {
		o[f] = v
	}
}

func (o Object) SetIfExists(f string, v interface{}) {
	if o == nil {
		return
	}
	if v != nil && v != "" {
		o[f] = v
	}
}

func (o Object) SetNotNilF(f string, vf func(f string) (interface{}, error)) error {
	if o == nil {
		return nil
	}
	v, err := vf(f)
	if err != nil {
		return err
	}
	if v != nil {
		o[f] = v
	}
	return nil
}

func (o Object) Clear() {
	if o == nil {
		return
	}
	for k := range o {
		delete(o, k)
	}
}

func (o Object) Remove(f string) {
	if o == nil {
		return
	}
	delete(o, f)
}

func (o Object) RemoveIf(pred func(f string, v interface{}) bool) {
	if o == nil {
		return
	}
	for f, v := range o {
		if !pred(f, v) {
			delete(o, f)
		}
	}
}

func (o Object) String() string {
	return jsonx.MarshalIndentString(o, "")
}
