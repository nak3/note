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

## graph (djacency matrix)
## graph (incidence matrix)
