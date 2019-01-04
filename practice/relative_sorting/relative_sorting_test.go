package relative_sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	a1  []int
	a2  []int
	exp []int
}{
	{
		[]int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8},
		[]int{2, 1, 8, 3},
		[]int{2, 2, 1, 1, 8, 8, 3, 5, 6, 7, 9},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.a1, tc.a2), "diff:%v", tc)
	}
}
