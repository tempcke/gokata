package main

import (
	"fmt"
	"testing"
)

var tt = []struct {
	names    []string
	greeting string
}{
	{[]string{"Bob"}, "Hello, Bob."},
	{[]string{}, "Hello, my friend."},
	{[]string{"JERRY"}, "HELLO JERRY!"},
	{[]string{"Jill", "Jane"}, "Hello, Jill and Jane."},
	{[]string{"Amy", "Brian", "Charlotte"}, "Hello, Amy, Brian, and Charlotte."},
	{[]string{"Amy", "BRIAN", "Charlotte"}, "Hello, Amy and Charlotte. AND HELLO BRIAN!"},
	{[]string{"Bob", "Charlie, Dianne"}, "Hello, Bob, Charlie, and Dianne."},
	{[]string{"Bob", "\"Charlie, Dianne\""}, "Hello, Bob and Charlie, Dianne."},
}

func TestGreeting(t *testing.T) {
	for i, tc := range tt {
		test := fmt.Sprintf("Requirment %d: %s", i, tc.greeting)
		t.Run(test, func(t *testing.T) {
			assertEqual(t, tc.greeting, greet(tc.names...))
		})
	}
}

func assertEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if got != want {
		t.Errorf(
			"Not Equal\nWant: %v\t%T\nGot:  %v\t%T",
			want, want,
			got, got,
		)
	}
}
