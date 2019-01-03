package transitive_closure

import (
	//"fmt"
	"testing"

	//	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	graph [][]int
	exp   [][]int
}{
	{
		[][]int{
			{1, 1, 0, 1},
			{0, 1, 1, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		},
		[][]int{{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, transitive_closure(tc.graph), "diff:%v", tc)
	}
}
