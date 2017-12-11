package randx

import (
	"math/rand"
	crand "crypto/rand"
	"fmt"
)

var (
	LowerLetters   = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters        = LowerLetters + UpperLetters
	Numbers        = "0123456789"
	LettersNumbers = Letters + Numbers
)

func Str(set string, n int) string {
	set1 := []rune(set)
	nSet := len(set1)
	r := make([]rune, n)
	for i := range r {
		r[i] = set1[rand.Intn(nSet)]
	}
	return string(r)
}
func RandomStr(n int) string {
	set1 := []rune(LettersNumbers)
	nSet := len(set1)
	r := make([]rune, n)
	for i := range r {
		r[i] = set1[rand.Intn(nSet)]
	}
	return string(r)
}

func InitToken(size int) string {
	b := make([]byte, size)
	crand.Read(b)
	return fmt.Sprintf("%x", b)
}
