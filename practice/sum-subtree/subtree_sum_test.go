package subtree_sum

import (
	"testing"

	"github.com/nak3/note/lib"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	target int
	root   func() *lib.Node
	exp    int
}{

	{
		7,
		func() *lib.Node { return lib.InsertLevelOrder([]int{5, -10, 3, 9, 8, -4, 7}, &lib.Node{}, 0, 7) },
		2,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, countSubtreesWithSumX(tc.target, tc.root()), "diff:%v", tc)
	}
}
