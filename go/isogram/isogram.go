package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	charList := make(map[rune]bool)
	for _, c := range word {
		if c != ' ' && c != '-' {
			_, present := charList[c]
			if present {
				return false
			}
			charList[c] = true
		}
	}
	return true

}
