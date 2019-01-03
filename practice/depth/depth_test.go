package depth

import (
	//"fmt"
	"testing"

	//	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	graph [][]int
	exp   int
}{

	[][]int{{5, 6}, {1, 3}, {2, 3}, {3, 4}, {1, 4}, {2, 5}, {3, 5}}, 4,
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, maxDepth(tc.graph), "diff:%v", tc)
	}
}
