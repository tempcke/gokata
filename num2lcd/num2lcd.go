package num2lcd

import (
	"strings"
)

func Num2Lcd(num int) string {
	return Scaled(num, 1, 1)
}

func Scaled(num, w, h int) string {
	digits := IntToSlice(num)

	var scaledDigits [][]string
	for i := len(digits)-1; i>=0; i-- {
		digit := digitRows(digits[i])
		scaledDigits = append(scaledDigits, scaleDigit(digit, w, h))
	}

	combined := combineRows(scaledDigits)
	return stringify(combined)
}

func IntToSlice(n int) []int {
	if n==0 { return []int{0} }
	return intToSliceIterator(n, []int{})
}

func intToSliceIterator(i int, s []int) []int {
	if i == 0 {
		return s
	}
	return intToSliceIterator(i/10, append(s, i%10))
}

func digitRows(n int) digit {
	return digits[n]
}

func scaleDigit(d digit, w, h int) []string {
	tallDigit := scaleDigitHeight(d[:], h)
	return scaleDigitWidth(tallDigit, w)
}

func scaleDigitWidth(rows []string, w int) []string {
	var result []string
	for _, row := range rows {
		result = append(result, scaleDigitRowWidth(row, w))
	}
	return result
}

func scaleDigitRowWidth(row string, w int) string {
	first := row[:1]
	mid := row[1:2]
	last := row[2:]
	return first + strings.Repeat(mid, w) + last
}

func scaleDigitHeight(rows []string, h int) []string {
	result := []string{rows[0]} // first element not repeated
	tail   := rows[1:]     // all but first element

	for _, row := range tail {
		result = append(result, scaleDigitRowHeight(row, h)...)
	}
	return result
}

func scaleDigitRowHeight(row string, h int) []string {
	first := row[:1]
	last  := row[2:]

	var result []string
	for j := 0; j < h-1; j++ {
		result = append(result, first + " " + last)
	}

	result = append(result, row)
	return result
}


func combineRows(digits [][]string) []string {
	result := make([]string, len(digits[0]))
	for _, digit := range digits {
		for i, row := range digit {
			result[i] += row
		}
	}
	return result
}

func stringify(rows []string) string {
	return strings.Join(rows, "\n") + "\n"
}

type digit [3]string

var digits = [...]digit{
	{
		" _ ",
		"| |",
		"|_|",
	}, {
		"   ",
		"  |",
		"  |",
	}, {
		" _ ",
		" _|",
		"|_ ",
	}, {
		" _ ",
		" _|",
		" _|",
	}, {
		"   ",
		"|_|",
		"  |",
	}, {
		" _ ",
		"|_ ",
		" _|",
	}, {
		" _ ",
		"|_ ",
		"|_|",
	}, {
		" _ ",
		"  |",
		"  |",
	}, {
		" _ ",
		"|_|",
		"|_|",
	}, {
		" _ ",
		"|_|",
		" _|",
	},
}
