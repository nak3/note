package trie_insert_and_search

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

func solve(base, search string) bool {
	data := strings.Split(base, " ")
	n := newTrie()
	for i := 0; i < len(data); i++ {
		n.insert(data[i])
	}
	if n.search(search) {
		return true
	}
	return false
}

func main() {
	s := "the a there answer any by bye their"
	search := "the"
	fmt.Printf("%+v\n", solve(s, search))
}
