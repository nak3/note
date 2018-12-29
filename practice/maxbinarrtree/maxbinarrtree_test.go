package practice

import (
	//"fmt"
	"testing"

	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums []int
	exp  []int
}{

	{
		[]int{3, 2, 1, 6, 0, 5},
		[]int{1, 2, 3, 0, 5, 6},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, lib.Postorder(constructMaximumBinaryTree(tc.nums)), "diff:%v", tc)
	}
}
