package check_for_balanced_tree

import (
	"testing"

	"github.com/nak3/note/lib"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	base []int
	exp  bool
}{
	{
		[]int{1, 10, 39, 50},
		true,
	},
	{
		[]int{1, 10, -99999, 50},
		false,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(lib.InsertLevelOrder(tc.base, &lib.Node{}, 0, len(tc.base))), "diff:%v", tc)
	}
}
