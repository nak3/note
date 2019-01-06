package vertical_width_of_a_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base []int
	exp  int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, -99999, -99999, -99999, -99999, -99999, 8, -99999, 9},
		6,
	},
	{
		[]int{1, 2, 3},
		3,
	},
	{
		[]int{10, 20, 30, 40, 60, 90},
		4,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		r := lib.InsertLevelOrder(tc.base, &lib.Node{}, 0, len(tc.base))
		ast.Equal(tc.exp, verticalWidth(r), "diff:%v", tc)
	}
}
