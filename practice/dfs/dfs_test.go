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
	[][]int{{1, 2}, {1, 3}, {1, 4}, {3, 5}},
	[]int{1, 2, 3, 5, 4},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, dfs_stack(tc.graph), "diff:%v", tc)
	}
	for _, tc := range tcs {
		ast.Equal(tc.exp, dfs_recursion(1, tc.graph), "diff:%v", tc)
	}

}
