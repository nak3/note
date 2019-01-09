package xor_cipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base string
	exp  string
}{
	{
		"A1D0A1D",
		"abcd",
	},
	{
		"653CAE8DA8EDB426052",
		"636f646572",
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base), "diff:%v", tc)
	}
}
