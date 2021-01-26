package main

import (
	"fmt"
	"strings"
	"unicode"
)

const greeting = "Hello, my friend."

func greet(names ...string) string {
	if len(names) == 0 {
		return greeting
	}
	norm, yell := segregate(names...)
	if len(norm) == 0 {
		return formatYelling(yell...)
	}
	if len(yell) == 0 {
		return formatNormal(norm...)
	}
	return fmt.Sprintf(
		"%s AND %s",
		formatNormal(norm...),
		formatYelling(yell...))
}

func segregate(names ...string) (normal, yelling []string) {
	for _, s := range names {
		for _, name := range splitNames(s) {
			if isUpper(name) {
				yelling = append(yelling, name)
				continue
			}
			normal = append(normal, name)
		}
	}
	return normal, yelling
}

func splitNames(s string) []string {
	if s[0] == '"' {
		return []string{s[1 : len(s)-1]}
	}
	return strings.Split(s, ", ")
}

func isUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func formatNormal(names ...string) string {
	nameStr := nameList(names...)
	return fmt.Sprintf("Hello, %s.", nameStr)
}

func formatYelling(names ...string) string {
	nameStr := nameList(names...)
	return fmt.Sprintf("HELLO %s!", nameStr)
}

func nameList(names ...string) string {
	if len(names) < 3 {
		return strings.Join(names, " and ")
	}
	n := len(names) - 1
	return strings.Join(names[:n], ", ") + ", and " + names[n]
}
