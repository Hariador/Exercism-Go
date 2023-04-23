package cryptosquare

import (
	"strings"
	"unicode"
)

func Encode(pt string) string {
	var sb strings.Builder
	for _, v := range pt {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			sb.WriteRune(unicode.ToLower(v))
		}
	}
	normalized := sb.String()
	if len(normalized) == 0 {
		return ""
	}
	c, r := getSizes(len(normalized))
	groups := groupStrings(normalized, c, r)
	sb.Reset()
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			sb.WriteRune(rune(groups[j][i]))
		}
		sb.WriteRune(' ')
	}
	output := sb.String()
	return output[:len(output)-1]
}

func getSizes(length int) (int, int) {
	c := 1
	r := 1
	for c*r < length {
		c++
		if c*r >= length {
			break
		}
		r++
		if c*r >= length {
			break
		}
	}
	return c, r
}
func groupStrings(normalized string, c, r int) []string {
	var groups = make([]string, r)
	for i := 0; i < r-1; i++ {
		groups[i] = normalized[i*c : i*c+c]
	}

	d := c*r - len(normalized)
	lastRow := normalized[r*c-c:]
	if d != 0 {

		for d > 0 {
			d--
			lastRow = lastRow + " "
		}
	}
	groups[r-1] = lastRow
	return groups

}
