package countPath

import "fmt"

type Graph struct {
	n     int
	edges [][]int
}

func newGraph(n int) *Graph {
	return &Graph{n, make([][]int, n)}
}

func (g *Graph) addNewEdge(s, d int) {
	g.edges[s] = append(g.edges[s], d)
}

func (g *Graph) count(s, d int, cnt *int, visit []bool) *int {
	if !visit[s] {
		visit[s] = true
	}
	if s == d {
		*cnt++
	} else {
		for _, v := range g.edges[s] {
			if visit[v] {
				continue
			}
			g.count(v, d, cnt, visit)
		}
	}
	visit[s] = false
	return cnt
}

func countPath(s, d, v int, edges [][]int) int {
	g := newGraph(v)

	for i := 0; i < len(edges); i++ {
		g.addNewEdge(edges[i][0], edges[i][1])

	}
	cnt := 0
	visit := make([]bool, v)
	ans := g.count(s, d, &cnt, visit)
	return *ans
}

// func main() {
// 	fmt.Printf("%+v\n", countPath(2, 3, 4, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 0}, {2, 1}, {1, 3}}))
// }
