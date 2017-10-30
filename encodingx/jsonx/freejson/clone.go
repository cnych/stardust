package freejson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Cloneable interface {
	FJClone() (interface{}, error)
}

func Clone(v interface{}) (interface{}, error) {
	var cloneValue func(interface{}) (interface{}, error)
	cloneValue = func(v interface{}) (interface{}, error) {
		if v == nil {
			return nil, nil
		}

		switch v1 := v.(type) {
		case Object:
			o1 := make(Object, len(v1))
			for f, fv := range v1 {
				fv1, err := cloneValue(fv)
				if err != nil {
					return nil, err
				}
				o1[f] = fv1
			}
			return o1, nil

		case Array:
			a1 := make(Array, 0, len(v1))
			for _, e := range v1 {
				e1, err := cloneValue(e)
				if err != nil {
					return nil, err
				}
				a1 = append(a1, e1)
			}
			return a1, nil

		case string, bool, json.Number:
			return v1, nil

		case time.Time:
			t1 := v1
			return t1, nil

		case Cloneable:
			return v1.FJClone()
		}
		return nil, fmt.Errorf("Not support value type for clone (%s)", reflect.TypeOf(v).Name())
	}
	return cloneValue(v)
}

func MustClone(v interface{}) interface{} {
	v1, err := Clone(v)
	if err != nil {
		panic(err)
	}
	return v1
}
