package array_subset_of_another_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	a   []int
	b   []int
	exp bool
}{
	{
		[]int{11, 1, 13, 21, 3, 7},
		[]int{11, 3, 7, 13},
		true,
	},
	{
		[]int{11, 1, 13, 21, 3, 7},
		[]int{11, 3, 7, 2},
		false,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.a, tc.b), "diff:%v", tc)
	}
}
