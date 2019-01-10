package find_the_closest_prime_number_to_a_given_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base int
	exp  int
}{
	{14, 13},
	{4, 3},
	{101, 101},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base), "diff:%v", tc)
	}
}
