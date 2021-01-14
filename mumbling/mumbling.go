package mumbleing

import "strings"

func mumbleLetters(input string) string {
	result := ""
	if len(input) == 0 {
		return result
	}

	var tokens []string
	for i, c := range strings.ToLower(input) {
		char := string(c)
		tok := strings.Repeat(char, i+1)
		tokens = append(tokens, strings.Title(tok))
	}
	return strings.Join(tokens, "-")
}