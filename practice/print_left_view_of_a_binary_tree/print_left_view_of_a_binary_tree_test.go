package print_left_view_of_a_binary_tree

import (
	"testing"

	"github.com/nak3/note/lib"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	base []int
	exp  []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 4},
	},
	{
		[]int{4, 5, 2, -99999, -99999, 3, 1, -99999, -99999, -99999, -99999, 6, 7},
		[]int{4, 5, 3, 6},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		root := &lib.Node{}
		i := 0
		n := len(tc.base)
		node := lib.InsertLevelOrder(tc.base, root, i, n)
		ast.Equal(tc.exp, solve(node), "diff:%v", tc)
	}
}
