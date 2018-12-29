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
