package subtree_sum

import (
	"fmt"

	"github.com/nak3/note/lib"
)

// util ...
func util(target, sum, l, r, ans int, node *lib.Node) (int, int) {
	if node == nil {
		return sum, ans
	}
	left, tmp1 := util(target, sum, l, r, ans, node.Left)
	right, tmp2 := util(target, sum, l, r, ans, node.Right)
	ans += tmp1 + tmp2
	if target == left+right+node.Val {
		ans++
	}

	fmt.Printf("%v %v %+v %d\n", left, right, node.Val, ans) // output for debug
	return left + right + node.Val, ans
}

func countSubtreesWithSumX(target int, root *lib.Node) int {
	if root == nil {
		return 0
	}
	_, ans := util(target, 0, 0, 0, 0, root)
	return ans
}

func main() {
	left := &lib.Node{3, nil, nil}
	right := &lib.Node{1, nil, nil}
	root := &lib.Node{2, left, right}
	ans := countSubtreesWithSumX(3, root)
	fmt.Printf("%+v\n", ans) // output for debug
}
