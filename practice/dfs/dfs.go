package dfs

//package main

import "fmt"

type Graph struct {
	edges map[int][]int
}

func newGraph() *Graph {
	return &Graph{map[int][]int{}}
}

func (g *Graph) addNewEdge(s, d int) {
	g.edges[s] = append(g.edges[s], d)
}

// dfs by stack
func dfs_stack(start int, edges [][]int) []int {
	g := newGraph()

	for i := 0; i < len(edges); i++ {
		g.addNewEdge(edges[i][0], edges[i][1])

	}

	visited := [101]bool{}
	q := []int{start}
	ans := []int{}
	for {
		if len(q) == 0 {
			break
		}
		tmp := q[len(q)-1]
		ans = append(ans, tmp)
		visited[tmp] = true
		q = q[:len(q)-1]

		for i := len(g.edges[tmp]) - 1; i >= 0; i-- {
			if !visited[g.edges[tmp][i]] {
				q = append(q, g.edges[tmp][i]) // output for debug
			}

		}
	}
	return ans
}

func (g *Graph) dfs_util(start int, visited [101]bool, ans *[]int) *[]int {
	if visited[start] {
		return ans
	}

	*ans = append(*ans, start)
	visited[start] = true

	for _, v := range g.edges[start] {
		g.dfs_util(v, visited, ans)
	}

	return ans
}

// dfs by recursion
func dfs_recursion(start int, edges [][]int) *[]int {
	g := newGraph()

	for i := 0; i < len(edges); i++ {
		g.addNewEdge(edges[i][0], edges[i][1])

	}
	visited := [101]bool{}
	ans := []int{}
	return g.dfs_util(start, visited, &ans)
}

// func main() {
// 	e := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 5}}
// 	ans := dfs_stack(1, e)
// 	fmt.Printf("DFS by stack\n") // output for debug
// 	for _, v := range ans {
// 		fmt.Printf("%+v ", v) // output for debug

// 	}
// 	fmt.Printf("\n")                 // output for debug
// 	fmt.Printf("DFS by recursion\n") // output for debug
// 	ans2 := dfs_recursion(1, e)
// 	for _, v := range *ans2 {
// 		fmt.Printf("%+v ", v) // output for debug

// 	}
// 	fmt.Printf("\n") // output for debug

// }
