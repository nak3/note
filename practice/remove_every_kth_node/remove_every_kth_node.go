//package main
package remove_every_kth_node

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) insert(newVal int) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &Node{newVal, nil}
}

func (n *Node) delete(key int) *Node {
	i := 1
	prev := &Node{}
	for n.Next != nil {
		fmt.Printf("%v %+v\n", i, n) // output for debug
		if i == key {
			prev.Next = n.Next
			n = n.Next
			i = 0
		} else {
			prev = n
			n = n.Next
		}
		i++
	}
	if i == key && n != nil {
		prev.Next = nil
	}

	return nil
}

func (n *Node) solve(key int) []int {
	ans := []int{}
	n.delete(key)
	for n != nil {
		ans = append(ans, n.Val)
		n = n.Next
	}
	return ans
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//	data := []int{1, 2, 3, 4}
	key := 2

	n := &Node{data[0], nil}
	for i := 1; i < len(data); i++ {
		n.insert(data[i])
	}
	fmt.Printf("%+v\n", n.solve(key))
	fmt.Printf("end\n") // output for debug
}
