package josephus_problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	n, k int
	exp  int
}{
	{3, 2, 3},
	{5, 3, 4},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.k), "diff:%v", tc)
	}
}
