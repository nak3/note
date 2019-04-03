package josephus_problem

// package REPLACE

import (
	"fmt"
	//	"github.com/nak3/note/lib"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

// solve ...
func solve(n, k int) int {
	head := &Node{1, nil, nil}
	p := head
	for i := 2; i < n; i++ {
		new := &Node{i, p, nil}
		p.next = new
		p = new
	}
	p.next = &Node{n, p, head}
	head.prev = p.next

	for {
		if head.prev == head && head == head.next {
			break
		}
		for i := 1; i < k; i++ {
			head = head.next
		}
		head.prev.next = head.next
		head.next.prev = head.prev
		head = head.next
	}
	return head.val
}

func main() {
	fmt.Printf("%+v\n", solve(3, 2)) // output for debug
	fmt.Printf("%+v\n", solve(5, 3)) // output for debug
}
