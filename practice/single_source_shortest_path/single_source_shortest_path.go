package single_source_shortest_path

import (
	"container/heap"
	"fmt"
)

// Priority queue library //
type IntHeap []pair

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].total < h[j].total } // Replace < with > for descending order
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pair))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type edge struct {
	to   int
	cost int
}

type pair struct {
	total int
	next  int
}

// solve ...
func solve(v, e, r int, data [][]int) []int {
	edges := make([][]*edge, v)
	dist := make([]int, v)
	done := make([]bool, v)
	for i := 0; i < e; i++ {
		s, t, cost := data[i][0], data[i][1], data[i][2]
		edges[s] = append(edges[s], &edge{t, cost})
	}

	queue := &IntHeap{}
	heap.Init(queue)

	INFTY := 1 << 31
	for i := 0; i < v; i++ {
		dist[i] = INFTY
	}

	heap.Push(queue, pair{0, r})

	for queue.Len() > 0 {
		item := heap.Pop(queue).(pair)
		if done[item.next] {
			continue
		}
		done[item.next] = true
		dist[item.next] = item.total
		for _, edge := range edges[item.next] {
			heap.Push(queue, pair{
				edge.cost + item.total,
				edge.to})
		}
	}
	ans := make([]int, v)
	for i, node := range dist {
		ans[i] = node
	}
	return ans
}

func main() {
	v := 4
	e := 6
	r := 1
	data := [][]int{
		{0, 1, 1},
		{0, 2, 4},
		{2, 0, 1},
		{1, 2, 2},
		{3, 1, 1},
		{3, 2, 5},
	}

	fmt.Printf("%+v\n", solve(v, e, r, data))
}
