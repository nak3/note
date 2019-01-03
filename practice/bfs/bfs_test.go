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
	[]int{1, 2, 3, 4, 6, 5},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, bfs(tc.graph), "diff:%v", tc)
	}
}
