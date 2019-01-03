package maxdiffancestor

import (
	"fmt"
	"math"

	"github.com/nak3/note/lib"
)

func util(mi, mx, ans int, node *lib.Node) int {
	if node == nil {
		return ans
	}
	tmp_mx := int(math.Abs(float64(node.Val - mi)))
	if tmp_mx > ans {
		ans = tmp_mx
	}
	if node.Val < mi {
		mi = node.Val
	}

	tmp_mi := int(math.Abs(float64(node.Val - mi)))
	if tmp_mi > ans {
		ans = tmp_mi
	}
	if node.Val < mx {
		mx = node.Val
	}

	l := util(mi, mx, ans, node.Left)
	r := util(mi, mx, ans, node.Right)

	if l > r {
		return l
	}
	return r
}

func maxDiff(root *lib.Node) int {
	if root == nil {
		return 0
	}
	mi, mx := root.Val, root.Val
	return util(mi, mx, 0, root)
}

func main() {
	left := &lib.Node{1, nil, nil}
	right := &lib.Node{2, nil, nil}
	root := &lib.Node{5, left, right}
	ans := maxDiff(root)
	fmt.Printf("%+v\n", ans) // output for debug
}
