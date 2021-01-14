package num2lcd

import (
	"gotest.tools/assert"
	"testing"
)

type req struct {
	number   int
	expected string
}

var singleDigitReqs = []req{
	{0, " _ \n| |\n|_|\n"},
	{1, "   \n  |\n  |\n"},
	{2, " _ \n _|\n|_ \n"},
	{3, " _ \n _|\n _|\n"},
	{4, "   \n|_|\n  |\n"},
	{5, " _ \n|_ \n _|\n"},
	{6, " _ \n|_ \n|_|\n"},
	{7, " _ \n  |\n  |\n"},
	{8, " _ \n|_|\n|_|\n"},
	{9, " _ \n|_|\n _|\n"},
}

func TestZero(t *testing.T) {
	assert.Equal(t, singleDigitReqs[0].expected, Num2Lcd(singleDigitReqs[0].number))
}

func TestSingleDigits(t *testing.T) {
	for _, req := range singleDigitReqs {
		assert.Equal(t, req.expected, Num2Lcd(req.number))
	}
}

func TestTwelve(t *testing.T) {
	expected := "    _ \n  | _|\n  ||_ \n"
	assert.Equal(t, expected, Num2Lcd(12))
}

func TestOneThreeSevenFive(t *testing.T) {
	expected := "    _  _  _ \n  | _|  ||_ \n  | _|  | _|\n"
	assert.Equal(t, expected, Num2Lcd(1375))
}

func TestSingleDigitHeightTwo(t *testing.T) {
	expected := " _ \n  |\n _|\n|  \n|_ \n"
	actual := Scaled(2, 1, 2)
	assert.Equal(t, expected, actual)
}

func TestSingleDigitHeightThree(t *testing.T) {
	expected := " _ \n  |\n  |\n _|\n|  \n|  \n|_ \n"
	actual := Scaled(2, 1, 3)
	assert.Equal(t, expected, actual)
}

func TestCombineRows(t *testing.T) {
	expected := []string{
		"    _  _ ",
		"  | _| _|",
		"  ||_  _|",
	}
	assert.Equal(t, stringify(expected), Scaled(123,1,1))
}

func TestCombineScaled(t *testing.T) {
	expected := []string{
		"     __  __ ",
		"   |   |   |",
		"   |   |   |",
		"   | __| __|",
		"   ||      |",
		"   ||      |",
		"   ||__  __|",
	}
	assert.Equal(t, stringify(expected), Scaled(123,2,3))
}