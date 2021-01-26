# Mumblng Kata

https://learn.madetech.com/katas/mumbling/

The goal of this kata is to:
1. implement the `mumble_letters(input string) string` method which takes a string as input and returns a formatted output string.
2. The output string contains sequences of repeating letters with each letter repeated a number of times based on its position in the input string i.e. the 3rd letter in the string is repeated 3 times.
3. Each sequence of repeated letters is separated with a hyphen(-)
4. the first letter of each sequence is capitalised.

The following examples illustrate the mumble_letters() method:

mumble_letters("")
=> ""

mumble_letters("a")
=> "A"

mumble_letters("abC")
=> "A-Bb-Ccc"

mumble_letters("aBCd")
=> "A-Bb-Ccc-Dddd"

mumble_letters("QWERTY")
=> "Q-Ww-Eee-Rrrr-Ttttt-Yyyyyy"