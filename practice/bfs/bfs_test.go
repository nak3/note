package bfs

import (
	//"fmt"
	"testing"

	//	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	graph [][]int
	exp   []int
}{
	[][]int{{0, 1}, {0, 2}, {0, 3}, {2, 0}, {2, 1}, {1, 3}},
	[]int{1, 2, 3, 4, 6, 5},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, bfs(tc.graph), "diff:%v", tc)
	}
}
