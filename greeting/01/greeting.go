package main

import (
	"regexp"
	"strings"
)

const defaultName = "my friend"

var isQuotedRegEx = regexp.MustCompile(`^".+"$`)

func Greet(nameList ...string) string {
	names := buildNames(nameList)
	return names.Greet()
}

func buildNames(names []string) Names {
	names = parseNameList(names)
	if len(names) == 1 && names[0] == "" {
		names[0] = defaultName
	}
	var normal []string
	var shout []string
	for _, name := range names {
		if name == strings.ToUpper(name) {
			shout = append(shout, name)
			continue
		}
		normal = append(normal, name)
	}
	return Names{normal, shout}
}

func parseNameList(input []string) []string {
	var names []string
	for _, val := range input {
		if !isQuoted(val) && strings.Contains(val, ",") {
			for _, name := range strings.Split(val, ",") {
				names = append(names, strings.Trim(name, " "))
			}
			continue
		}
		names = append(names, strings.Trim(val, "\""))
	}
	return names
}

func isQuoted(name string) bool {
	return isQuotedRegEx.MatchString(name)
}
