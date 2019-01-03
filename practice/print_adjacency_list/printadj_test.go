package printadj

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
		[][]int{{0, 1}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {3, 4}},
		[][]int{{1, 4}, {0, 2, 3, 4}, {1, 3}, {1, 2, 4}, {0, 1, 3}},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, arrayAdj(tc.graph), "diff:%v", tc)
	}
}
