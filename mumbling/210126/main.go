package main

import (
	"strings"
	"unicode"
)

func mumbleLetters(in string) string {
	var sb strings.Builder
	for i, r := range in {
		if i > 0 {
			sb.WriteRune('-')
		}
		A, a := unicode.ToUpper(r), unicode.ToLower(r)
		sb.WriteRune(A)
		for j := 0; j < i; j++ {
			sb.WriteRune(a)
		}
	}
	return sb.String()
}
