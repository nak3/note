package smallest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	base []int
	exp  []string
}{
	{
		[]int{6, 3, 9, 8, 10, 2, 1, 15, 7},
		[]string{"7", "6", "10", "9", "15", "3", "2", "_", "8"},
	},
	{
		[]int{13, 6, 7, 12},
		[]string{"_", "7", "12", "13"},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base), "diff:%v", tc)
	}
}
