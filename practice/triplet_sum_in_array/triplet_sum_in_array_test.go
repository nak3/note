package triplet_sum_in_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	target int
	data   []int
	exp    int
}{
	{
		13,
		[]int{1, 4, 45, 6, 10, 8},
		1,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.data, tc.target), "diff:%v", tc)
	}
}
