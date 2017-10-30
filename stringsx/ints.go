package stringsx

import "strconv"

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Itob(i int) bool {
	if i != 0 {
		return true
	}
	return false
}

func Strtoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return 0
	}
	return i
}

func Strtof(s string) float64 {
	f, e := strconv.ParseFloat(s, 32)
	if e != nil {
		return 0
	}
	return f
}
