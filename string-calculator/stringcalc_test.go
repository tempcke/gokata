package string_calculator

import (
	"gotest.tools/assert"
	"testing"
)

type requirement struct {
	desc     string
	input    string
	expected int
}

var requirements = []requirement{
	{
		"An empty string returns zero",
		"",
		0,
	},
	{
		"A single number returns the value",
		"42",
		42,
	},
	{
		"Two numbers, comma delimited, returns the sum",
		"2,3",
		5,
	},
	{
		"Two numbers, newline delimited, returns the sum",
		"3\n4",
		7,
	},
	{
		"Three numbers, delimited either way, returns the sum",
		"3\n4\n5",
		12,
	},
	{
		"Three numbers, delimited either way, returns the sum",
		"3,4,5",
		12,
	},
	{
		"Numbers greater than 1000 are ignored",
		"1,1000,1001,2",
		1003,
	},
	{
		"Numbers greater than 1000 are ignored",
		"1001",
		0,
	},
	{
		"A single char delimiter can be defined on the first line (e.g. //# for a ‘#’ as the delimiter)",
		"//#\n3#4",
		7,
	},
	{
		"A multi char delimiter can be defined on the first line (e.g. //[###] for ‘###’ as the delimiter)",
		"//[###]\n3###4",
		7,
	},
	{
		"Many single or multi-char delimiters can be defined (each wrapped in square brackets)",
		"//[##][%%]\n3##4%%5",
		12,
	},
}

func TestStringCalculator(t *testing.T) {
	for _, req := range requirements {
		t.Run(req.desc, func(t *testing.T) {
			checkRequirement(t, req)
		})
	}
}

func TestErrorHandling(t *testing.T) {
	assertErrNegNums(t, "-42")
	assertErrNegNums(t, "4,-2")
	assertErrNegNums(t, "4\n-2")
}

func assertErrNegNums(t *testing.T, input string) {
	t.Helper()
	n, e := Resolve(input)
	assert.Equal(t, ErrNegNums, e, "\nExpected error due to negative number\n\tInput: %s\n\tResult: %d", input, n)
}

func checkRequirement(t *testing.T, req requirement) {
	result, e := Resolve(req.input)
	if e != nil {
		t.Errorf("\nUnexpected error\n\tInput: %s\n\tError: %s\n", req.input, e.Error())
	}
	if req.expected != result {
		t.Errorf("\nWant  %d\nGot   %d\n", req.expected, result)
	}
}
