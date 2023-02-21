package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	if err != nil {
		panic(fmt.Sprintf("Regex didnt compile: %v", err))
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~\*=-]*>`)
	if err != nil {
		panic(fmt.Sprintf("Regex didn't compile: %v", err))
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[[:digit:]]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	userNameRe := regexp.MustCompile(`User +[[:alnum:]]+ `)
	nameRe := regexp.MustCompile(`User +`)
	var temp []string
	for _, line := range lines {
		if userNameRe.MatchString(line) {
			subLine := userNameRe.FindStringSubmatch(line)
			name := nameRe.ReplaceAllString(subLine[0], "")
			logLine := "[USR] " + name + line
			temp = append(temp, logLine)
		} else {
			temp = append(temp, line)
		}

	}
	return temp
}
