package main

import (
	"fmt"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestStringInts(t *testing.T) {
	tt := []struct {
		in  string
		out []string
	}{
		{"1,2,3", []string{"1", "2", "3"}},
		{"1\n2\n3", []string{"1", "2", "3"}},
		{"//#\n1#2#3", []string{"1", "2", "3"}},
		{"//[$$]\n1$$2$$3", []string{"1", "2", "3"}},
		{"//[$$][#]\n1#2$$3", []string{"1", "2", "3"}},
		{"//#\n", []string{}},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			assert.Equal(t,
				strings.Join(tc.out, ","),
				strings.Join(stringInts(tc.in), ","),
			)
		})
	}
}

func TestStringCalc(t *testing.T) {
	var tt = []struct {
		input  string
		output int
		test   string
	}{
		{"", 0, "1. An empty string returns zero"},
		{"1", 1, "2. A single number returns the value"},
		{"42", 42, "2. A single number returns the value"},
		{"1,2", 3, "3. Two numbers, comma delimited, returns the sum"},
		{"4\n5", 9, "4. Two numbers, newline delimited, returns the sum"},
		{"1,2\n3", 6, "5. Three numbers, delimited either way, returns the sum"},
		{"1,1000,1001", 1001, "7. Numbers greater than 1000 are ignored"},
		{"//#\n1#2#3", 6, "8. A single char delimiter can be defined on first line via //# where # is the delim char"},
		{"//[$$]\n1$$2$$3", 6, "9. A multi char delimiter can be defined on first line via //[###] where ### is the delim substring"},
		{"//[$$][#]\n1#2$$3", 6, "10. Support many delimiters each wrapped in square brackets"},
		{"//#\n", 0, "1. An empty string returns zero"},
	}
	for _, tc := range tt {
		name := fmt.Sprintf("%s - %s", tc.test, tc.input)
		t.Run(name, func(t *testing.T) {
			n, err := strcalc(tc.input)
			if err != nil {
				t.Error(err)
			}
			assertEqual(t, tc.output, n)
		})
	}
}

func TestStringCalcErrors(t *testing.T) {
	var tt = []string{
		"a",
		"1,b",
		"a,1",
		"a,b",
		"5,-2", // 6. Negative numbers cause error
	}
	for _, tc := range tt {
		name := fmt.Sprintf("Expect Error From: %s", tc)
		t.Run(name, func(t *testing.T) {
			_, err := strcalc(tc)
			if err == nil {
				t.Errorf("Expected error from: %s", tc)
			}
		})
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf(
			"Values are not equal\nWant %v\t%T\nGot  %v\t%T",
			expected, expected,
			actual, actual,
		)
	}
}
