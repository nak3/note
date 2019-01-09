package minimum_platforms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	arrive    []int
	departure []int
	exp       int
}{
	{
		[]int{900, 940, 950, 1100, 1500, 1800},
		[]int{910, 1200, 1120, 1130, 1900, 2000},
		3,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.arrive, tc.departure), "diff:%v", tc)
	}
}
