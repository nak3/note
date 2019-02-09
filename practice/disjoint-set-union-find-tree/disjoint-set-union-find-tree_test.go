package disjoint_set_union_find_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	n    int
	q    int
	data [][]int
	exp  []int
}{
	{
		5,
		12,
		[][]int{
			{0, 1, 4},
			{0, 2, 3},
			{1, 1, 2},
			{1, 3, 4},
			{1, 1, 4},
			{1, 3, 2},
			{0, 1, 3},
			{1, 2, 4},
			{1, 3, 0},
			{0, 0, 4},
			{1, 0, 2},
			{1, 3, 0},
		},
		[]int{0, 0, 1, 1, 1, 0, 1, 1},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.q, tc.data), "diff:%v", tc)
	}
}
