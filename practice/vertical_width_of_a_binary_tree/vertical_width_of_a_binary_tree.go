package vertical_width_of_a_binary_tree

//package main

import (
	"fmt"
	"github.com/nak3/note/lib"
)

// util ...
func util(r *lib.Node, miL, mxR int) (int, int) {
	if r == nil {
		return miL, mxR
	}
	var tmpLL, tmpLR, tmpRL, tmpRR int
	if r.Left != nil {
		tmpLL, tmpLR = util(r.Left, miL-1, mxR-1)
	}
	if r.Right != nil {
		tmpRL, tmpRR = util(r.Right, miL+1, mxR+1)
	}

	if tmpRL < tmpLL {
		tmpLL = tmpRL
	}
	if tmpLR > tmpRR {
		tmpRR = tmpLR
	}

	if tmpLL < miL {
		miL = tmpLL
	}
	if tmpRR > mxR {
		mxR = tmpRR
	}
	return miL, mxR
}

func verticalWidth(root *lib.Node) int {
	l, r := util(root, 0, 0)
	return r - l + 1
}

func main() {
	nodes := []int{1, 2, 3, 4, 5, 6, 7, -99999, -99999, -99999, -99999, -99999, 8, -99999, 9}
	//nodes := []int{1, 2, 3}
	r := lib.InsertLevelOrder(nodes, &lib.Node{}, 0, len(nodes))
	fmt.Printf("%+v\n", verticalWidth(r)) // output for debug
}
