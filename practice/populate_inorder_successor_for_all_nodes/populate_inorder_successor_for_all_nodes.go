package populate_inorder_successor_for_all_nodes

//package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nak3/note/lib"
)

func solve(root *lib.Node, l, r *lib.Node, ans map[int][]int) map[int][]int {
	if root == nil {
		return ans
	}
	if root.Left != nil {
		ans = solve(root.Left, root, nil, ans)
	}
	if l != nil {
		fmt.Printf("%v->%v\n", root.Val, l.Val) // output for debug
		ans[root.Val] = append(ans[root.Val], l.Val)
	}
	if r != nil {
		fmt.Printf("%v->%v\n", r.Val, root.Val) // output for debug
		ans[r.Val] = append(ans[r.Val], root.Val)
		if root.Left == nil && root.Right == nil {
			ans[root.Val] = append(ans[root.Val], -1)
		}
	}
	if root.Right != nil {
		ans = solve(root.Right, nil, root, ans)
	}
	return ans
}

func populateNext(root *lib.Node) string {
	mp := map[int][]int{}
	out := solve(root, nil, nil, mp)
	tmp := []int{}
	for k, _ := range out {
		tmp = append(tmp, k)
	}
	ret := ""
	sort.Ints(tmp)
	for i := 0; i < len(tmp); i++ {
		val := out[tmp[i]]
		sort.Ints(val)
		s := strconv.Itoa(val[0])
		for i := 1; i < len(val); i++ {
			s += "->" + strconv.Itoa(val[i])
		}
		ret += strconv.Itoa(tmp[i]) + "->" + s
		ret += " "
	}
	ret = strings.TrimSpace(ret)
	return ret
}

func main() {
	arr := []int{10, 8, 12, 3, 13}
	root := &lib.Node{}
	i := 0
	n := len(arr)
	tree := lib.InsertLevelOrder(arr, root, i, n)
	ans := populateNext(tree)
	fmt.Printf("%+v\n", ans) // output for debug
}
