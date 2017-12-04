package randx

import (
	"math/rand"
	"reflect"
)

type W struct {
	W int         `json:"w" bson:"w"`
	V interface{} `json:"v" bson:"v"`
}

func ChoiceStr(choices ...string) string {
	n := len(choices)
	if n == 0 {
		return ""
	}
	return choices[rand.Intn(n)]
}

func ChoiceInt(choices ...int) int {
	n := len(choices)
	if n == 0 {
		return 0
	}
	return choices[rand.Intn(n)]
}

func ChoiceInt64(choices ...int64) int64 {
	n := len(choices)
	if n == 0 {
		return 0
	}
	return choices[rand.Intn(n)]
}

func Choice(choices ...interface{}) interface{} {
	n := len(choices)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return choices[0]
	}
	return choices[rand.Intn(n)]
}

func ChoiceSlice(choicesSlice interface{}) interface{} {
	if choicesSlice == 0 {
		return nil
	}
	v := reflect.ValueOf(choicesSlice)
	k := v.Kind()
	if k == reflect.Slice || k == reflect.Array {
		n := v.Len()
		if n == 0 {
			return nil
		}
		return v.Index(rand.Intn(n)).Interface()
	} else {
		return nil
	}
}

func ChoiceW(choices ...W) interface{} {
	n := len(choices)
	if n == 0 {
		return nil
	}
	if n == 1 {
		first := choices[0]
		if first.W > 0 {
			return first.V
		} else {
			return nil
		}
	}
	var sum, upto int64 = 0, 0
	for _, w := range choices {
		if w.W > 0 {
			sum += int64(w.W)
		}
	}
	r := Float64Between(0.0, float64(sum))
	for _, w := range choices {
		ww := w.W
		if ww < 0 {
			ww = 0
		}
		if float64(upto)+float64(ww) > r {
			return w.V
		}
		upto += int64(w.W)
	}
	return nil
}
