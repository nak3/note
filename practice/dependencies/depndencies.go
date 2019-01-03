package dependencies

import "fmt"

type Graph struct {
	n     int
	edges [][]int
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) count(V int) int {
	cnt := 0
	for v := 0; v < V; v++ {
		cnt += len(g.edges[v])
	}
	return cnt
}

func NewGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: make([][]int, n),
	}
	return g
}

func depend(v, e int, edges []int) int {
	g := NewGraph(v)
	for i := 0; i < len(edges); i += 2 {
		g.addEdge(edges[i], edges[i+1])
	}
	return g.count(v)
}

// func main() {
// 	tmp := depend(4, 4, []int{0, 2, 0, 3, 1, 3, 2, 3})
// 	fmt.Printf("%+v\n", tmp) // output for debug

// 	tmp = depend(4, 3, []int{0, 2, 0, 1, 0, 3})
// 	fmt.Printf("%+v\n", tmp) // output for debug
// }
