package single_source_shortest_path

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	v    int
	e    int
	r    int
	data [][]int
	exp  []int
}{
	{
		4,
		6,
		1,
		[][]int{
			{0, 1, 1},
			{0, 2, 4},
			{2, 0, 1},
			{1, 2, 2},
			{3, 1, 1},
			{3, 2, 5},
		},
		[]int{3, 0, 2, 1 << 31},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.v, tc.e, tc.r, tc.data), "diff:%v", tc)
	}
}
