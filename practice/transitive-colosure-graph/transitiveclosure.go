package transitive_closure

import "fmt"

func transitive_closure(graph [][]int) [][]int {
	v := len(graph)
	ans := graph
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				if ans[i][k] != 0 && ans[k][j] != 0 {
					ans[i][j] = 1
				}
			}
		}
	}
	return ans
}

func main() {
	e := [][]int{
		{1, 1, 0, 1},
		{0, 1, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 0, 1},
	}
	ans := transitive_closure(e)
	fmt.Printf("%+v\n", ans) // output for debug
}
