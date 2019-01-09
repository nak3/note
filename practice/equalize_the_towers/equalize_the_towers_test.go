package equalize_the_towers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	height []int
	cost   []int
	exp    int
}{
	{
		[]int{1, 2, 3},
		[]int{10, 100, 1000},
		120,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.height, tc.cost), "diff:%v", tc)
	}
}
