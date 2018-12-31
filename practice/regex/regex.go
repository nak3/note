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

// https://www.hackerrank.com/challenges/matching-whitespace-non-whitespace-character/problem
func regex_whitespace(str string) bool {
	valid := regexp.MustCompile(`\S{2}\s\S{2}\s\S{2}`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/matching-word-non-word/problem
func regex_word(str string) bool {
	valid := regexp.MustCompile(`\w{3}\W\w{10}\W\w{3}`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/matching-start-end/problem
func regex_startend(str string) bool {
	valid := regexp.MustCompile(`^\d\w{4}\.$`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/matching-specific-characters/problem
func regex_specific(str string) bool {
	valid := regexp.MustCompile(`^[1-3][0-2][xs0][30Aa][xsu][.,]$`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/excluding-specific-characters/problem
func regex_notspecific(str string) bool {
	valid := regexp.MustCompile(`\D[^aeiou][^bcDF]\S[^AEIOU][^.,]$`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/matching-range-of-characters/problem
func regex_range(str string) bool {
	valid := regexp.MustCompile(`[a-z][1-9][^a-z][^A-Z][A-Z]$`)
	return valid.MatchString(str)
}

// https://www.hackerrank.com/challenges/matching-x-repetitions/problem
func regex_repetitions(str string) bool {
	valid := regexp.MustCompile(`[a-zA-Z+]{40}[\s02468+]{5}$`)
	return valid.MatchString(str)
}
