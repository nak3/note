package bfs

//package

//import "fmt"

type Graph struct {
	edges map[int][]int
}

func newGraph() *Graph {
	return &Graph{map[int][]int{}}
}

func (g *Graph) addNewEdge(s, d int) {
	g.edges[s] = append(g.edges[s], d)
}

func bfs(start int, edges [][]int) []int {
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
		tmp := q[0]
		ans = append(ans, q[0])
		visited[q[0]] = true
		q = q[1:]

		for i := 0; i < len(g.edges[tmp]); i++ {
			if !visited[g.edges[tmp][i]] {
				q = append(q, g.edges[tmp][i]) // output for debug
			}

		}
	}
	return ans
}

// func main() {
// 	e := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 5}, {1, 6}}
// 	ans := bfs(1, e)
// 	for _, v := range ans {
// 		fmt.Printf("%+v ", v) // output for debug

// 	}
// 	fmt.Printf("\n") // output for debug

// }
