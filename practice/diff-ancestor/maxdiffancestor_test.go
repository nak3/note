package maxdiffancestor

import (
	"testing"

	"github.com/nak3/note/lib"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	root func() *lib.Node
	exp  int
}{

	{
		func() *lib.Node { return lib.InsertLevelOrder([]int{5, 1, 2}, &lib.Node{}, 0, 3) },
		4,
	},
	{
		func() *lib.Node { return lib.InsertLevelOrder([]int{0, 1, 2, 3, 4, 5, 6}, &lib.Node{}, 0, 7) },
		6,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		node := tc.root()
		ast.Equal(tc.exp, maxDiff(node), "diff:%v", tc)
	}
}
