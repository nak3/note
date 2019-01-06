# basic data structure

Basic data strctures in Go. Keep it simple as much as possible.

## stack

~~~go
package main

import "fmt"

func main() {
	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)

	n := len(stack) - 1
	fmt.Println(stack[n]) // top()

	stack = stack[:n] // pop()

	n = len(stack) - 1
	fmt.Println(stack[n]) // top()
}
~~~

## queue

~~~go
package main

import "fmt"

func main() {
	queue := make([]int, 0)

	queue = append(queue, 1) // push()
	queue = append(queue, 2) // push()

	n := queue[0] // front()

	fmt.Println(n)

	queue = queue[1:] // pop()

	n = queue[0] // front()
	fmt.Println(n)
}
~~~

## priority queue

This is [an example in golang.org](https://golang.org/pkg/container/heap/)

```go
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
```

## singly linkedlist 

```go
package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func (n *Node) add(val int) {
	newNode := Node{val, nil}
	for true {
		if n.next == nil {
			n.next = &newNode
			break
		} else {
			n = n.next
		}
	}
}

func main() {
	root := Node{0, nil}
	root.add(1)
	root.add(2)
	root.add(3)

	n := &root
	for true {
		if n == nil {
			break
		}
		fmt.Printf("%+v\n", n.val)
		n = n.next
	}
}
```

## doubly linkedlist

```go
package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

func (n *Node) add(val int) {
	newNode := Node{val, nil, nil}
	for true {
		if n.next == nil {
			n.next = &newNode
			newNode.prev = n
			break
		} else {
			n = n.next
		}
	}
}

func main() {
	root := Node{0, nil, nil}
	root.add(1)
	root.add(2)
	root.add(3)

	n := &root
	for true {
		if n == nil {
			break
		}
		fmt.Printf("%v %+v %+v\n", n.val, n.prev, n.next)
		n = n.next
	}
}
```

## tree

```go
package main

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func preorder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Printf("%v ", t.data)
	preorder(t.left)
	preorder(t.right)

}

func postorder(t *Tree) {
	if t == nil {
		return
	}
	postorder(t.left)
	postorder(t.right)
	fmt.Printf("%v ", t.data)
}

func inorder(t *Tree) {
	if t == nil {
		return
	}
	inorder(t.left)
	fmt.Printf("%v ", t.data)
	inorder(t.right)
}

/*
      0
    1   2
  3  4   5
*/

func main() {
	five := &Tree{5, nil, nil}
	four := &Tree{4, nil, nil}
	three := &Tree{3, nil, nil}
	two := &Tree{2, nil, five}
	one := &Tree{1, three, four}
	zero := &Tree{0, one, two}
	fmt.Printf("preorder: ")
	preorder(zero)
	fmt.Printf("\n")
	fmt.Printf("postorder: ")
	postorder(zero)
	fmt.Printf("\n")
	fmt.Printf("inorder: ")
	inorder(zero)
}
```

## graph (by slice)

```go
package main

import "fmt"

func addEdge(adj [][]int, u, v int) [][]int {
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
	return adj
}

func printGraph(adj [][]int, V int) {
	for v := 0; v < V; v++ {
		fmt.Printf("adj of %+v is: ", v) // output for debug
		for _, x := range adj[v] {
			fmt.Printf("%+v ", x) // output for debug
		}
		fmt.Printf("\n") // output for debug
	}
}

func main() {
	/*
	     3
	     |
	   2-1-4
	     |
	     5
	*/
	V := 5
	adj := make([][]int, V+1)
	addEdge(adj, 1, 2)
	addEdge(adj, 1, 3)
	addEdge(adj, 1, 4)
	addEdge(adj, 1, 5)
	printGraph(adj, V)
}
```


## graph (adjacency matrix)

```go
package main

import "fmt"

func addEdge(adj [][]int, u, v int) [][]int {
	adj[u][v] = 1
	adj[v][u] = 1
	return adj
}

func printGraph(adj [][]int, V int) {
	fmt.Printf("   ") // output for debug
	for v := 0; v < V; v++ {
		fmt.Printf("%+v ", v) // output for debug
	}
	fmt.Printf("\n") // output for debug

	for v := 0; v < V; v++ {
		fmt.Printf("%+v: ", v) // output for debug
		for i, x := range adj[v] {
			if i == 0 {
				continue
			}
			fmt.Printf("%+v ", x) // output for debug
		}
		fmt.Printf("\n") // output for debug
	}
}

func main() {
	/*
	     3
	     |
	   2-1-4
	     |
	     5
	*/
	V := 5
	adj := make([][]int, V+1)
	for i := 0; i < V+1; i++ {
		adj[i] = make([]int, V+1)
	}
	addEdge(adj, 1, 2)
	addEdge(adj, 1, 3)
	addEdge(adj, 1, 4)
	addEdge(adj, 1, 5)
	printGraph(adj, V)
}
```
## undirected graph

```go
package main

import "fmt"

type Graph struct {
	n     int
	edges [][]int
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

func (g *Graph) printGraph(V int) {
	for v := 0; v < V; v++ {
		fmt.Printf("from %+v to: ", v+1) // output for debug
		for _, x := range g.edges[v] {
			fmt.Printf("%+v ", x+1) // output for debug
		}
		fmt.Printf("\n") // output for debug
	}
}

func NewGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: make([][]int, n),
	}
	return g
}

func main() {
	/*
		             4
		             |
			   2-1-3
	*/
	V := 4
	g := NewGraph(V)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.printGraph(V)
}
```

## directed graph

```go
package main

import "fmt"

type Graph struct {
	n     int
	edges [][]int
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) printGraph(V int) {
	for v := 0; v < V; v++ {
		fmt.Printf("from %+v to: ", v) // output for debug
		for _, x := range g.edges[v] {
			fmt.Printf("%+v ", x) // output for debug
		}
		fmt.Printf("\n") // output for debug
	}
}

func NewGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: make([][]int, n),
	}
	return g
}

func main() {
	/*
		              4
		              ^
			   2<-1->3
	*/
	V := 4
	g := NewGraph(V)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.printGraph(V)
}
```

## directed graph (by map)

```go
package main

import "fmt"

type Graph struct {
	n     int
	edges map[int][]int
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
//	g.edges[v] = append(g.edges[v], u)
}

func (g *Graph) printGraph(V int) {
	for k, v := range g.edges {
		fmt.Printf("from %+v to: ", k) // output for debug
		for i := 0; i < len(v); i++ {
			fmt.Printf("%+v ", v[i]) // output for debug
		}
		fmt.Printf("\n") // output for debug
	}
}

func NewGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: map[int][]int{},
	}
	return g
}

func main() {
	/*
		             4
		             |
			   2-1-3
	*/
	V := 4
	g := NewGraph(V)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.printGraph(V)
}
```
