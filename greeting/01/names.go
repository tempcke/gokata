package main

import (
	"fmt"
	"strings"
)

type Names struct {
	normal []string
	shout  []string
}

func (names Names) Greet() string {
	normalGreet := names.normalGreet()
	shoutGreet  := names.shoutGreet()

	if names.hasShout() {
		if names.hasNormal() {
			return fmt.Sprintf("%s AND %s", normalGreet, shoutGreet)
		}
		return shoutGreet
	}
	return normalGreet
}

func (names Names) list(list []string) string {
	if len(list) == 2 {
		return strings.Join(list, " and ")
	}
	var result string
	for i, name := range list {
		switch i {
		case 0: result = name
		case len(list) - 1: result = fmt.Sprintf("%s, and %s", result, name)
		default: result = fmt.Sprintf("%s, %s", result, name)
		}
	}
	return result
}

func (names Names) hasNormal() bool {
	return len(names.normal) >= 1
}

func (names Names) hasShout() bool {
	return len(names.shout) >= 1
}

func (names Names) normalGreet() string {
	if !names.hasNormal() {
		return ""
	}
	return fmt.Sprintf("Hello, %s.", names.list(names.normal))
}

func (names Names) shoutGreet() string {
	if !names.hasShout() {
		return ""
	}
	return fmt.Sprintf("HELLO %s!", names.list(names.shout))
}