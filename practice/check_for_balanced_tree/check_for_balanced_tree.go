package check_for_balanced_tree

//package main

import (
	"fmt"
	"math"

	"github.com/nak3/note/lib"
)

// hight ...
func height(root *lib.Node) int {
	if root == nil {
		return 0
	}
	mr := height(root.Right)
	ml := height(root.Left)
	if mr > ml {
		return 1 + mr
	}
	return 1 + ml
}

func solve(root *lib.Node) bool {
	if root == nil {
		return true
	}
	heightR := height(root.Right)
	heightL := height(root.Left)

	if math.Abs(float64(heightR-heightL)) > 1 {
		return false
	}
	if solve(root.Right) && solve(root.Left) {
		return true
	}
	return false
}

func main() {
	nodes := []int{1, 10, 39, 50}
	//	nodes := []int{1, 10, -100, 50}
	root := &lib.Node{}
	n := lib.InsertLevelOrder(nodes, root, 0, len(nodes))
	fmt.Printf("%+v\n", solve(n)) // output for debug
}
