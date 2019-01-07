package diameter_of_binary_tree

//package main

import (
	"fmt"

	"github.com/nak3/note/lib"
)

// height ...
func height(node *lib.Node, ans *int) int {
	if node == nil {
		return 0
	}
	heightLeft := height(node.Left, ans)
	heightRight := height(node.Right, ans)
	if heightLeft+heightRight+1 > *ans {
		*ans = heightLeft + heightRight + 1
	}
	if heightLeft > heightRight {
		return heightLeft + 1
	}
	return heightRight + 1
}

// solve ...
func solve(node *lib.Node) int {
	if node == nil {
		return 0
	}
	ans := 0
	height(node, &ans)
	return ans
}

func main() {
	nodes := []int{10, 20, 30, 40, 60}
	//	nodes := []int{1, 2, 3}

	r := lib.InsertLevelOrder(nodes, &lib.Node{}, 0, len(nodes))
	fmt.Printf("%+v\n", solve(r)) // output for debug

}
