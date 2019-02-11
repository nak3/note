package all_pairs_shortest_path

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	v      int
	e      int
	data   [][]int
	exp    [][]int64
	expErr error
}{
	{
		4,
		6,
		[][]int{
			{0, 1, 1},
			{0, 2, 5},
			{1, 2, 2},
			{1, 3, 4},
			{2, 3, 1},
			{3, 2, 7},
		},
		[][]int64{
			{0, 1, 3, 4},
			{4294967296, 0, 2, 3},
			{4294967296, 4294967296, 0, 1},
			{4294967296, 4294967296, 7, 0},
		},
		nil,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ans, err := solve(tc.v, tc.e, tc.data)
		ast.Equal(tc.exp, ans, "diff:%v", tc)
		ast.Equal(tc.expErr, err, "diff:%v", tc)
	}
}
