package printadj

//package main

import (
	"fmt"
	"sort"
)

type Graph struct {
	edges map[int][]int
}

func newGraph() *Graph {
	return &Graph{map[int][]int{}}
}

func (g *Graph) addNewEdge(s, d int) {
	g.edges[s] = append(g.edges[s], d)
	g.edges[d] = append(g.edges[d], s)
}

func (g *Graph) printAdj() {
	keys := []int{}
	for k, _ := range g.edges {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Printf("%+v ", v) // output for debug
		for _, vv := range g.edges[v] {
			fmt.Printf("-> %+v ", vv) // output for debug
		}
		fmt.Printf("\n") // output for debug
	}
}

func (g *Graph) arrayAdj() [][]int {
	ans := make([][]int, len(g.edges))
	keys := []int{}
	for k, _ := range g.edges {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		for _, vv := range g.edges[v] {
			ans[v] = append(ans[v], vv)
		}
	}
	return ans
}

func arrayAdj(e [][]int) [][]int {
	g := newGraph()
	for i := 0; i < len(e); i++ {
		g.addNewEdge(e[i][0], e[i][1])
	}
	return g.arrayAdj()
}

// func main() {
// 	g := newGraph()
// 	e := [][]int{{0, 1}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {3, 4}}

// 	for i := 0; i < len(e); i++ {
// 		g.addNewEdge(e[i][0], e[i][1])
// 	}

// 	g.printAdj()
// }
