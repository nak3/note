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
