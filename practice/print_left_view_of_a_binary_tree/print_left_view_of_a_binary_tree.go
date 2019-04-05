package print_left_view_of_a_binary_tree

// package REPLACE

import (
	"fmt"
	"github.com/nak3/note/lib"
)

// solve ...
func solve(root *lib.Node) []int {
	tmp := []*lib.Node{root}

	ans := []int{}
	for len(tmp) != 0 {
		length := len(tmp)
		ans = append(ans, tmp[0].Val)
		for i := 0; i < length; i++ {
			if tmp[i].Left != nil {
				tmp = append(tmp, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				tmp = append(tmp, tmp[i].Right)
			}
		}
		tmp = tmp[length:]
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	root := &lib.Node{}
	i := 0
	n := len(arr)
	tmp := lib.InsertLevelOrder(arr, root, i, n)
	fmt.Printf("%+v\n", solve(tmp)) // output for debug
}
