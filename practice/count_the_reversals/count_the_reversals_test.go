package count_the_reversals

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base string
	exp  int
}{
	{"}{{}}{{{", 3},
	{"{{}}}}", 1},
	{"{{}{{{}{{}}{{", -1},
	{"{{{{}}}}", 0},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base), "diff:%v", tc)
	}
}
