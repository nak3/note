package flood_fill_algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	data  [][]int
	color int
	place []int
	exp   [][]int
}{
	{
		[][]int{
			{0, 1, 1},
			{0, 1, 1},
			{1, 1, 0},
			{1, 2, 3},
		},
		5,
		[]int{0, 1},
		[][]int{{0, 5, 5}, {0, 5, 5}, {5, 5, 0}, {5, 2, 3}},
	},
	{
		[][]int{{1, 1, 1, 1}},
		8,
		[]int{0, 1},
		[][]int{{8, 8, 8, 8}},
	},
	{
		[][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 3, 2, 4}},
		9,
		[]int{0, 2},
		[][]int{{1, 2, 9, 4}, {1, 2, 9, 4}, {1, 2, 9, 4}, {1, 3, 2, 4}},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.place, tc.color, tc.data), "diff:%v", tc)
	}
}
