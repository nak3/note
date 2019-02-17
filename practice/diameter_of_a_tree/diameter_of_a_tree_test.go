package diameter_of_a_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	n    int
	data [][]int
	exp  int
}{
	{
		4,
		[][]int{
			{0, 1, 2},
			{1, 2, 1},
			{1, 3, 3},
		},
		5,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.data), "diff:%v", tc)
	}
}
