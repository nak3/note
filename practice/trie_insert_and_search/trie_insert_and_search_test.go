package trie_insert_and_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"github.com/nak3/note/lib"
)

var tcs = []struct {
	base   string
	search string
	exp    bool
}{
	{
		"the a there answer any by bye their",
		"the",
		true,
	},
	{
		"the a there answer any by bye their",
		"a",
		true,
	},
	{
		"the a there answer any by bye their",
		"they",
		false,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.exp, solve(tc.base, tc.search), "diff:%v", tc)
	}
}
