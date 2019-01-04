package save_winterfell

//package main

import (
	"fmt"
)

type Graph struct {
	n     int
	edges [][]int
}

func newGraph(n int) *Graph {
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = make([]int, n)
	}

	return &Graph{n, e}
}

func (g *Graph) solve(edge int, v, d *[8]int, n int) int {
	mx := 0
	if v[edge] != 0 {
		return d[edge]
	}
	for i := 0; i < n; i++ {
		if g.edges[edge][i] > 0 {
			tmp := g.solve(i, v, d, n) + g.edges[edge][i]
			if tmp > mx {
				mx = tmp
			}
		}
	}
	d[edge] = mx
	v[edge] = 1
	return mx
}

func solve(e [][]int, n int) int {
	g := newGraph(n)
	for i := 0; i < len(e); i++ {
		// operates -1 as vertex starts 1.
		g.edges[e[i][0]-1][e[i][1]-1] = e[i][2]
	}
	fmt.Printf("%+v\n", g) // output for debug

	visited := [8]int{}
	distance := [8]int{}
	mx := 0
	for i := 0; i < n; i++ {
		tmp := g.solve(i, &visited, &distance, n)
		if tmp > mx {
			mx = tmp
		}
	}
	return mx
}

func main() {
	n := 8
	edges := [][]int{
		{1, 5, 3},
		{3, 5, 4},
		{5, 4, 2},
		{2, 4, 5},
		{4, 7, 6},
		{7, 8, 2},
		{7, 6, 1},
	}
	fmt.Printf("%+v\n", solve(edges, n))

}
