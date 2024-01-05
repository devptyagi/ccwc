package charactercount

import "unicode/utf8"

// CountCharacters counts the number of characters in the given input, considering UTF-8 encoding.
func CountCharacters(input []byte) int {
	return utf8.RuneCount(input)
}
