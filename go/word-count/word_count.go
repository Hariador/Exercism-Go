package wordcount

import (
	"strings"
	"unicode"
)

const (
	ScanningDigits int = 0
	ScanningWord       = 1
	ScanningNew        = 2
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var tokens []string
	phrase = strings.ToLower(phrase)
	state := ScanningNew
	token := ""
	for _, char := range phrase {
		if state == ScanningNew {
			if unicode.IsLetter(char) {
				state = ScanningWord
				token = token + string(char)
			}
			if unicode.IsDigit(char) {
				state = ScanningDigits
				token = token + string(char)
			}
		} else {
			if state == ScanningDigits {
				if unicode.IsDigit(char) {
					token = token + string(char)
				} else {
					state = ScanningNew
					tokens = append(tokens, token)
					token = ""
				}
			}
			if state == ScanningWord {
				if unicode.IsLetter(char) || char == '\'' {
					token = token + string(char)
				} else {
					state = ScanningNew
					tokens = append(tokens, rTrim(token))
					token = ""
				}

			}
		}
	}
	if token != "" {
		tokens = append(tokens, rTrim(token))
	}
	temp := make(map[string]int)
	for _, v := range tokens {
		if _, ok := temp[v]; ok {
			temp[v] = temp[v] + 1
		} else {
			temp[v] = 1
		}
	}

	return temp
}

func rTrim(token string) string {
	if token[len(token)-1] == '\'' {
		return token[:len(token)-1]
	}
	return token
}
