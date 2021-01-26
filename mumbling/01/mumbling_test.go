package mumbleing

import "testing"

type requirement struct {
	input    string
	expected string
}

var requirements = []requirement{
	{"", ""},
	{"a", "A"},
	{"ab", "A-Bb"},
	{"abC", "A-Bb-Ccc"},
	{"aBCd", "A-Bb-Ccc-Dddd"},
	{"QWERTY", "Q-Ww-Eee-Rrrr-Ttttt-Yyyyyy"},
}

func TestMumble(t *testing.T) {
	for _, req := range requirements {
		t.Run(req.input+" to "+req.expected, func(t *testing.T) {
			checkRequirement(t, req)
		})
	}
}

func checkRequirement(t *testing.T, req requirement) {
	result := mumbleLetters(req.input)
	if req.expected != result {
		t.Errorf(
			"\nInput %s\nWant  %s\nGot   %s\n",
			req.input,
			req.expected,
			result)
	}
}

func BenchmarkMumbleLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mumbleLetters("abcdefghijklmnopqrstuvwxyz")
	}
}
