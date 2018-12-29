package practice

import (
	//"fmt"
	"testing"

	"../../lib"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str string
	exp int
}{
	{"123", 123},
	{"12a", -1},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, atoi(tc.str), "diff:%v", tc)
	}
}
