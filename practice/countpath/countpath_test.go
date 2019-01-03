package countpath

import (
	//"fmt"
	"testing"

	//	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	start    int
	dest     int
	vertices int
	graph    [][]int
	exp      int
}{
	2,
	3,
	4,
	[][]int{{0, 1}, {0, 2}, {0, 3}, {2, 0}, {2, 1}, {1, 3}},
	3,
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, countPath(tc.graph), "diff:%v", tc)
	}
}
