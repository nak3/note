package depth

//package main

import "fmt"

type Graph struct {
	edges map[int][]int
}

func NewGraph() *Graph {
	g := &Graph{
		edges: map[int][]int{},
	}
	return g
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) depth(k int, mp map[int]int) int {
	val, ok := mp[k]
	if ok {
		return val
	}
	maxLeng := 0
	for _, val := range g.edges[k] {
		tmp := g.depth(val, mp)
		if maxLeng < tmp+1 {
			maxLeng = tmp + 1
		}
	}
	mp[k] = maxLeng
	return maxLeng
}

func maxDepth(edge [][]int) int {
	g := NewGraph()
	for i := 0; i < len(edge); i++ {
		g.addEdge(edge[i][0], edge[i][1])
	}
	mp := map[int]int{}
	mx := 0
	for k, _ := range g.edges {
		tmp := g.depth(k, mp)
		if tmp > mx {
			mx = tmp
		}
	}
	return mx
}

// func main() {
// 	tmp := maxDepth([][]int{{5, 6}, {1, 3}, {2, 3}, {3, 4}, {1, 4}, {2, 5}, {3, 5}})
// 	fmt.Printf("%+v\n", tmp) // output for debug
// }
