package paths_to_reach_origin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	n   int
	m   int
	exp int
}{
	{3, 2, 10},
	{3, 6, 84},
	{3, 0, 1},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.n, tc.m), "diff:%v", tc)
	}
}
