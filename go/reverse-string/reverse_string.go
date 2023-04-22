package reverse

func Reverse(input string) string {
	if len(input) <= 1 {
		return ""
	}
	var temp []rune
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		temp = append(temp, runes[i])
	}
	return string(temp)
}
