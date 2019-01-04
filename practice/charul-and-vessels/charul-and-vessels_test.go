package charul_and_vessels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	items  []int
	target int
	exp    bool
}{
	{
		[]int{6, 3, 4, 2, 1},
		20,
		true,
	},
	{
		[]int{2, 4, 3},
		15,
		true,
	},
	{
		[]int{8, 3, 12, 6},
		10,
		false,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, dp(tc.target, tc.items), "diff:%v", tc)
	}
}
