package diameter_of_binary_tree

import (
	"testing"

	"github.com/nak3/note/lib"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	base []int
	exp  int
}{
	{
		[]int{10, 20, 30, 40, 60},
		4,
	},
	{
		[]int{1, 2, 3},
		3,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		r := lib.InsertLevelOrder(tc.base, &lib.Node{}, 0, len(tc.base))
		ast.Equal(tc.exp, solve(r), "diff:%v", tc)
	}
}
