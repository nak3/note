package reverse_first_k_elements_of_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	n   int
	k   int
	arr []int
	exp []int
}{
	{
		5,
		3,
		[]int{1, 2, 3, 4, 5},
		[]int{3, 2, 1, 4, 5},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.k, tc.arr), "diff:%v", tc)
	}
}
