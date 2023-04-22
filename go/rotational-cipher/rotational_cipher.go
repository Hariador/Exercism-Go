package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {

	var output []int32
	runeStream := []rune(plain)
	var rot int32
	for _, v := range runeStream {
		if unicode.IsUpper(v) {
			rot = (v - 39) + int32(shiftKey)
			rot = rot%26 + 65
		} else if unicode.IsLower(v) {
			rot = (v - 71) + int32(shiftKey)
			rot = rot%26 + 97
		} else {
			rot = v
		}

		output = append(output, rot)
	}

	return string(output)
}
