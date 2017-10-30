package stringsx

import (
	"sort"
	"strings"
)

func SliceRemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func SliceIndexOf(a []string, s string) int {
	for i, elem := range a {
		if elem == s {
			return i
		}
	}
	return -1
}

func SliceContains(a []string, s string) bool {
	return SliceIndexOf(a, s) >= 0
}

func ISliceIndexOf(a []int, s int) int {
	for i, elem := range a {
		if elem == s {
			return i
		}
	}
	return -1
}

func I64SliceIndexOf(a []int64, s int64) int {
	for i, elem := range a {
		if elem == s {
			return i
		}
	}
	return -1
}

func ISliceContains(a []int, s int) bool {
	return ISliceIndexOf(a, s) >= 0
}

func I64SliceContains(a []int64, s int64) bool {
	return I64SliceIndexOf(a, s) >= 0
}

func SliceContainsCase(a []string, s string, caseInsensitive bool) bool {
	if caseInsensitive {
		s = strings.ToLower(s)
	}
	for _, v := range a {
		if caseInsensitive {
			v = strings.ToLower(v)
		}
		if s == v {
			return true
		}
	}
	return false
}

// transforms each item of a slice to lowercase.
func SliceToLower(a []string) {
	for index, value := range a {
		a[index] = strings.ToLower(value)
	}
}

func SliceUnique(a []string) []string {
	n := len(a)
	if n == 0 {
		return []string{}
	}
	set := map[string]int{}
	a1 := make([]string, 0, n)
	for _, elem := range a {
		_, ok := set[elem]
		if !ok {
			a1 = append(a1, elem)
			set[elem] = 1
		}
	}
	return a1
}

func SliceInt64Unique(a []int64) []int64 {
	n := len(a)
	if n == 0 {
		return []int64{}
	}
	set := map[int64]int{}
	a1 := make([]int64, 0, n)
	for _, elem := range a {
		_, ok := set[elem]
		if !ok {
			a1 = append(a1, elem)
			set[elem] = 1
		}
	}
	return a1
}

func SliceFlatten(al ...[]string) []string {
	a1 := make([]string, 0, 4)
	for _, a := range al {
		for _, elem := range a {
			a1 = append(a1, elem)
		}
	}
	return a1
}

func SliceClone(a []string) []string {
	if a == nil {
		return nil
	}
	n := len(a)
	if n == 0 {
		return []string{}
	}
	a1 := make([]string, n, n)
	for i, elem := range a {
		a1[i] = elem
	}
	return a1
}

func SliceMap(a []string, f func(s string) string) []string {
	n := len(a)
	if n == 0 {
		return []string{}
	}
	a1 := make([]string, 0, n)
	for _, elem := range a {
		a1 = append(a1, f(elem))
	}
	return a1
}

func SliceEqual(a1 []string, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func SliceAsSetEqual(a1 []string, a2 []string) bool {
	a1 = SliceUnique(a1)
	sort.Strings(a1)
	a2 = SliceUnique(a2)
	sort.Strings(a2)
	return SliceEqual(a1, a2)
}

func SliceFilter(a []string, target []string, filter func(s string) bool) []string {
	for _, s := range a {
		if filter(s) {
			target = append(target, s)
		}
	}
	return target
}
