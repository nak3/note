package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str string
	exp bool
}{
	{"abc.def.ghi.jkx", true},
	{"abc.def.ghi", false},
	{"1.1.1.1", false},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, regex(tc.str), "diff:%v", tc)
	}
}

var tcs2 = []struct {
	str string
	exp bool
}{
	{"12A34B5678", true},
	{"1234567890", false},
}

func Test_fn2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs2 {
		ast.Equal(tc.exp, regex2(tc.str), "diff:%v", tc)
	}
}

var tcs3 = []struct {
	str string
	exp bool
}{
	{"12 34 56", true},
	{"1234567890", false},
}

func Test_whitespace(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs3 {
		ast.Equal(tc.exp, regex_whitespace(tc.str), "diff:%v", tc)
	}
}

var tcs4 = []struct {
	str string
	exp bool
}{
	{"abc#defghijklm$nop", true},
	{"abcdefghijklmnopq", false},
}

func Test_word(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs4 {
		ast.Equal(tc.exp, regex_word(tc.str), "diff:%v", tc)
	}
}

var tcs5 = []struct {
	str string
	exp bool
}{
	{"1abcd.", true},
	{"abcde.", false},
	{"2abcde", false},
}

func Test_startend(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs5 {
		ast.Equal(tc.exp, regex_startend(tc.str), "diff:%v", tc)
	}
}

var tcs6 = []struct {
	str string
	exp bool
}{
	{"11x3x.", true},
	{"41x3x.", false},
	{"41x3x2.", false},
}

func Test_specific(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs6 {
		ast.Equal(tc.exp, regex_specific(tc.str), "diff:%v", tc)
	}
}

var tcs7 = []struct {
	str string
	exp bool
}{
	{"A1A1aO", true},
	{"think?", true},
	{"1hink?", false},
	{"taink?", false},
	{"thbnk?", false},
	{"thi\nk?", false},
	{"thinA?", false},
	{"thinA.", false},
}

func Test_notspecific(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs7 {
		ast.Equal(tc.exp, regex_notspecific(tc.str), "diff:%v", tc)
	}
}

var tcs8 = []struct {
	str string
	exp bool
}{
	{"a1AaA", true},
	{"A1AaA", false},
	{"aaAaA", false},
	{"a1aaA", false},
	{"a1AAA", false},
	{"a1aAa", false},
}

func Test_range(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs8 {
		ast.Equal(tc.exp, regex_range(tc.str), "diff:%v", tc)
	}
}

var tcs9 = []struct {
	str string
	exp bool
}{
	{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa24680", true},
	{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa24681", false},
	{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa24680", false},
	{"1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa24680", false},
}

func Test_repetitions(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs9 {
		ast.Equal(tc.exp, regex_repetitions(tc.str), "diff:%v", tc)
	}
}
