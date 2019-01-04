package save_winterfell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	data [][]int
	num  int
	exp  int
}{
	{
		[][]int{
			{1, 5, 3},
			{3, 5, 4},
			{5, 4, 2},
			{2, 4, 5},
			{4, 7, 6},
			{7, 8, 2},
			{7, 6, 1},
		},
		8,
		14,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.data, tc.num), "diff:%v", tc)
	}
}
