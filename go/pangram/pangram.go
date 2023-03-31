package pangram

import "strings"

func IsPangram(input string) bool {
	letters := initMap()
	input = strings.ToLower(input)
	for _, letter := range input {
		letters[letter] = true
		if foundAll(letters) {
			return true
		}
	}

	return false
}

func initMap() map[rune]bool {
	temp := map[rune]bool{'a': false, 'b': false, 'c': false, 'd': false, 'e': false, 'f': false, 'g': false, 'h': false, 'i': false, 'j': false,
		'k': false, 'l': false, 'm': false, 'n': false, 'o': false, 'p': false, 'q': false, 'r': false, 's': false, 't': false, 'u': false,
		'v': false, 'w': false, 'x': false, 'y': false, 'z': false}
	return temp
}
func foundAll(letters map[rune]bool) bool {
	for _, v := range letters {
		if !v {
			return false
		}
	}
	return true
}
