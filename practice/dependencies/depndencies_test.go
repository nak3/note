package dependencies

import (
	//"fmt"
	"testing"

	//	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	vertex int
	edge   int
	edges  []int
	exp    int
}{
	{4, 4, []int{0, 2, 0, 3, 1, 3, 2, 3}}, 4,
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, depend(tc.vertex, tc.edge, tc.edges), "diff:%v", tc)
	}
}
