package atbash

import (
	"unicode"
)

func Atbash(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	reverse := "zyxwvutsrqponmlkjihgfedcba"
	cypher := make(map[rune]rune)
	for k, v := range alphabet {
		cypher[v] = rune(reverse[k])
	}

	output := ""
	outcount := 0
	for _, v := range s {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			if outcount == 5 {
				outcount = 0
				output = output + " "
			}
		}
		if unicode.IsDigit(v) {
			output = output + string(v)
			outcount++
		} else if unicode.IsSpace(v) {
			//noop
		} else if unicode.IsLetter(v) {
			output = output + string(cypher[unicode.ToLower(v)])
			outcount++
		}

	}
	return output
}
