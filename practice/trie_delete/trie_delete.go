package trie_delete

//package main

import (
	"fmt"
	"strings"
	//	"github.com/nak3/note/lib"
)

type Trie struct {
	Node       map[string]*Trie
	EndOfValue bool
}

func newTrie() *Trie {
	return &Trie{map[string]*Trie{}, false}
}

func (t *Trie) search(s string) bool {
	node := t
	for i := 0; i < len(s); i++ {
		idx := string(s[i])
		if _, ok := node.Node[idx]; !ok {
			return false
		}
		node = node.Node[idx]
	}
	return node != nil && node.EndOfValue
}

func (t *Trie) insert(s string) {
	node := t
	for i := 0; i < len(s); i++ {
		idx := string(s[i])
		if _, ok := node.Node[idx]; !ok {
			node.Node[idx] = newTrie()
		}
		node = node.Node[idx]
	}
	node.EndOfValue = true
}

func (t *Trie) remove(key string, depth int) *Trie {
	if t == nil {
		return nil
	}
	if depth == len(key) {
		if t.EndOfValue {
			t.EndOfValue = false
		}
		if len(t.Node) == 0 {
			t = nil
		}
		return t
	}

	idx := string(key[depth])
	tmp := t.Node[idx].remove(key, depth+1)
	t.Node[idx] = tmp
	if len(t.Node) == 0 && t.EndOfValue == false {
		t = nil
	}
	return t
}

func solve(base, search, del string) [2]bool {
	data := strings.Split(base, " ")
	n := newTrie()
	for i := 0; i < len(data); i++ {
		n.insert(data[i])
	}
	ans := [2]bool{}
	ans[0] = n.search(search)
	n.remove(del, 0)
	ans[1] = n.search(search)
	return ans
}

func main() {
	s := "the a there answer any by bye their"
	search, del := "there", "the"

	ans := solve(s, search, del)
	fmt.Printf("%v %v\n", ans[0], ans[1])
}
