package main

import (
	"fmt"
	"testing"
)

func TestMumbling(t *testing.T) {
	tt := []struct {
		in, out string
	}{
		{"", ""},
		{"a", "A"},
		{"abc", "A-Bb-Ccc"},
		{"QWERTY", "Q-Ww-Eee-Rrrr-Ttttt-Yyyyyy"},
	}
	for i, tc := range tt {
		test := fmt.Sprintf("%d. %s => %s", i, tc.in, tc.out)
		t.Run(test, func(t *testing.T) {
			assertEqual(t, tc.out, mumbleLetters(tc.in))
		})
	}
}

func assertEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if got != want {
		t.Errorf(
			"Not Equal!\nWant: %v\t%T\nGot:  %v\t%T",
			want, want,
			got, got)
	}
}

func BenchmarkMumbleLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mumbleLetters("abcdefghijklmnopqrstuvwxyz")
	}
}
