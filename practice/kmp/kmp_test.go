package kmp

import (
	//"fmt"
	"testing"

	//	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	base    string
	pattern string
	exp     int
}{
	{
		"ABABDABACDABABCABAB",
		"ABABCABAB",
		10,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, KMPSearch(tc.base, tc.pattern), "diff:%v", tc)
	}
}
