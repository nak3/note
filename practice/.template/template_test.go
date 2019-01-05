package REPLACE

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base string
	exp  int
}{
	{
		"ABABDABACDABABCABAB",
		10,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, TODO(tc.base), "diff:%v", tc)
	}
}
