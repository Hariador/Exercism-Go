package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	chars := []rune(isbn)
	pos := 10
	sum := 0
	for _, v := range chars {
		if unicode.IsDigit(v) {
			number, _ := strconv.Atoi(string(v))
			sum = sum + number*pos
			pos--
		} else if v == 'X' {
			if pos != 1 {
				return false
			}
			sum = sum + 10
			pos--
		} else if v != '-' {
			return false
		}
	}
	if sum == 0 || pos != 0 {
		return false
	}
	return sum%11 == 0

}
