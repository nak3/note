package range_minimum_quer

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
		3,
		5,
		[][]int{
			{0, 0, 1},
			{0, 1, 2},
			{0, 2, 3},
			{1, 0, 2},
			{1, 1, 2},
		},
		[]int{1, 2},
	},
	{
		1,
		3,
		[][]int{
			{1, 0, 0},
			{0, 0, 5},
			{1, 0, 0},
		},
		[]int{2147483647, 5},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.q, tc.data), "diff:%v", tc)
	}
}
