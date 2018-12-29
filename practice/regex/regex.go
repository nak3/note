package regex

import (
	"regexp"
)

// https://www.hackerrank.com/challenges/matching-anything-but-new-line/problem
func regex(str string) bool {
	valid := regexp.MustCompile(`[a-z|A-Z]{3}\.[a-z|A-Z]{3}\.[a-z|A-Z]{3}\.[a-z|A-Z]{3}`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/matching-digits-non-digit-character/problem
func regex2(str string) bool {
	valid := regexp.MustCompile(`[\d]{2}[\D][\d]{2}[\D][\d]{4}`)
	return valid.MatchString(str)
}
