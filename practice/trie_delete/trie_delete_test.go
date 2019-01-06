package trie_delete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base   string
	search string
	del    string
	exp    [2]bool
}{
	{
		"the a there answer any by bye their",
		"the",
		"the",
		[2]bool{true, false},
	},
	{
		"the a there answer any by bye their",
		"the",
		"there",
		[2]bool{true, true},
	},
	{
		"the a there answer any by bye their",
		"there",
		"the",
		[2]bool{true, true},
	},
	{
		"the a there answer any by bye their",
		"there",
		"non-existing",
		[2]bool{true, true},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base, tc.search, tc.del), "diff:%v", tc)
	}
}
