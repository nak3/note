package weighted_union_find_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	n    int
	q    int
	data [][]int
	exp  []string
}{
	{
		5,
		6,
		[][]int{
			{0, 0, 2, 5},
			{0, 1, 2, 3},
			{1, 0, 1, 0},
			{1, 1, 3, 0},
			{0, 1, 4, 8},
			{1, 0, 4, 0},
		},
		[]string{"2", "?", "10"},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.q, tc.data), "diff:%v", tc)
	}
}
