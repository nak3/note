package populate_inorder_successor_for_all_nodes

import (
	"testing"

	"github.com/nak3/note/lib"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	base []int
	exp  string
}{
	{
		[]int{10, 8, 12, 3, 13},
		"3->8 8->10->13 10->12 12->-1 13->-1",
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	root := &lib.Node{}
	for _, tc := range tcs {
		ast.Equal(tc.exp, populateNext(lib.InsertLevelOrder(tc.base, root, 0, len(tc.base))), "diff:%v", tc)
	}
}
