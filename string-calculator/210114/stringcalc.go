package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func strcalc(input string) (int, error) {
	if input == "" {
		return 0, nil
	}

	var sum int
	for _, tok := range stringInts(input) {
		n, err := strIntVal(tok)
		if err != nil {
			return sum, err
		}
		sum += n
	}
	return sum, nil
}

func strIntVal(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, errors.New("all numbers must be greater than or equal to zero")
	}
	if n > 1000 {
		return 0, nil
	}
	return n, nil
}

var (
	multiCharDelimPattern  = regexp.MustCompile(`\/\/\[(.+)\]\n(.+)`)
	singleCharDelimPattern = regexp.MustCompile(`\/\/(.)\n(.+)`)
)

func stringInts(input string) []string {
	var delim = ","
	delims := []string{"\n"}
	if strings.HasPrefix(input, "//") {
		r := singleCharDelimPattern
		if input[2] == '[' {
			r = multiCharDelimPattern
		}
		m := r.FindStringSubmatch(input)
		if m == nil { // resolves where input has a defined delim but no string
			return []string{}
		}
		input = m[2]
		delims = strings.Split(m[1], "][")
	}
	for _, d := range delims {
		input = strings.Replace(input, d, delim, -1)
	}
	return strings.Split(input, delim)
}
