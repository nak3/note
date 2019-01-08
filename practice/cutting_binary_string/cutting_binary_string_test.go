package cutting_binary_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base string
	exp  int
}{
	{
		"101101101",
		3,
	},
	{
		"101101",
		2,
	},
	{
		"1",
		-1,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base), "diff:%v", tc)
	}
}
