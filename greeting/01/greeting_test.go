package main

import (
	"testing"
)

type requirement struct {
	desc     string
	inputs   []string
	expected string
}

var requirements = []requirement{
	{
		"Requirement 1, simple greeting",
		[]string{"Bob"},
		"Hello, Bob.",
	}, {
		"Requirement 2, handle null name",
		[]string{""},
		"Hello, my friend.",
	}, {
		"Requirement 3, shout for uppercase names",
		[]string{"JERRY"},
		"HELLO JERRY!",
	}, {
		"Requirement 4, two names",
		[]string{"Jill", "Jane"},
		"Hello, Jill and Jane.",
	}, {
		"Requirement 5, three or more names",
		[]string{"Amy", "Brian", "Charlotte"},
		"Hello, Amy, Brian, and Charlotte.",
	}, {
		"Req 6, separate normal and shouting names",
		[]string{"Amy", "BRIAN", "Charlotte"},
		"Hello, Amy and Charlotte. AND HELLO BRIAN!",
	}, {
		"Req 7, detect multiple names within each string arg",
		[]string{"Bob", "Charlie, Dianne"},
		"Hello, Bob, Charlie, and Dianne.",
	}, {
		"Req 8, detect multiple names within each string arg",
		[]string{"Bob", "\"Charlie, Dianne\""},
		"Hello, Bob and Charlie, Dianne.",
	},
}

func TestGreetings(t *testing.T) {
	for _, req := range requirements {
		t.Run(req.desc, func(t *testing.T) {
			checkRequirement(t, req)
		})
	}
}

func checkRequirement(t *testing.T, req requirement) {
	greeting := Greet(req.inputs...)
	if req.expected != greeting {
		t.Errorf("\nWant  %s\nGot   %s\n", req.expected, greeting)
	}
}
