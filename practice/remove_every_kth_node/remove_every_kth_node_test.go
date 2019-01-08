package remove_every_kth_node

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base []int
	key  int
	exp  []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		2,
		[]int{1, 3, 5, 7},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		// TODO: library
		n := &Node{tc.base[0], nil}
		for i := 1; i < len(tc.base); i++ {
			n.insert(tc.base[i])
		}
		ast.Equal(tc.exp, n.solve(tc.key), "diff:%v", tc)
	}
}
