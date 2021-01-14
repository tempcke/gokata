# The String Calculator Kata

This kata, made popular by Roy Osherove, comes with a precise set of steps to follow. The essence is a method that given a delimited string, returns the sum of the values. I’ve always preferred my kata to define the tests I will follow every time through the exercise, so here are the tests I use for this one:

1. An empty string returns zero
2. A single number returns the value
3. Two numbers, comma delimited, returns the sum
4. Two numbers, newline delimited, returns the sum
5. Three numbers, delimited either way, returns the sum
6. Negative numbers cause error
7. Numbers greater than 1000 are ignored
8. A single char delimiter can be defined on the first line (e.g. //# for a ‘#’ as the delimiter)
9. A multi char delimiter can be defined on the first line (e.g. //[###] for ‘###’ as the delimiter)
10. Many single or multi-char delimiters can be defined (each wrapped in square brackets)

