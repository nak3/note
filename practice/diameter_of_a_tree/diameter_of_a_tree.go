//package main

package diameter_of_a_tree

// package REPLACE

import (
	"fmt"
	//	"github.com/nak3/note/lib"
)

type Tree struct {
	nodes      []map[int]int
	done       map[int]bool
	mxDistance int
	mxNode     int
}

// dfs ...
func (t *Tree) dfs(s int, sum int) {
	t.done[s] = true
	next := t.nodes[s]
	for k, v := range next {
		if t.done[k] {
			continue
		}
		t.dfs(k, sum+v)
	}
	if sum > t.mxDistance {
		t.mxDistance = sum
		t.mxNode = s
	}
	return
}

func solve(n int, data [][]int) int {
	t := &Tree{}

	t.nodes = make([]map[int]int, n)
	for i := 0; i < n; i++ {
		t.nodes[i] = map[int]int{}
	}

	for i := 0; i < n-1; i++ {
		s := data[i][0]
		g := data[i][1]
		w := data[i][2]

		t.nodes[s][g] = w
		t.nodes[g][s] = w
	}
	t.done = make(map[int]bool)
	t.dfs(0, 0)
	t.done = make(map[int]bool)
	t.mxDistance = 0
	start := t.mxNode
	t.dfs(start, 0)
	return t.mxDistance
}

func main() {
	n := 4
	data := [][]int{
		{0, 1, 2},
		{1, 2, 1},
		{1, 3, 3},
	}

	t := &Tree{}

	t.nodes = make([]map[int]int, n)
	for i := 0; i < n; i++ {
		t.nodes[i] = map[int]int{}
	}

	for i := 0; i < n-1; i++ {
		s := data[i][0]
		g := data[i][1]
		w := data[i][2]

		t.nodes[s][g] = w
		t.nodes[g][s] = w
	}
	t.done = make(map[int]bool)
	t.dfs(0, 0)
	t.done = make(map[int]bool)
	start := t.mxNode
	t.dfs(start, 0)
	fmt.Printf("%+v\n", t.mxDistance) // output for debug
}
