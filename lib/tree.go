package lib

import (
//	"fmt"
)

// Node represents binary tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// InsertLevelOrder creates LevelOrderTree from array
// usage example:
//   func main() {
//        arr := []int{1, 2, 3, 4, 5}
//        root := &Node{}
//        i := 0
//        n := len(arr)
//        ans := InsertLevelOrder(arr, root, i, n)
//    }
// NOTE: node -99999 is magic number. It inserts nil Node
func InsertLevelOrder(arr []int, root *Node, i, n int) *Node {
	if i < n {
		tmp := &Node{Val: arr[i]}
		if arr[i] == -99999 {
			tmp = nil
		} else {
			root = tmp
			root.Left = InsertLevelOrder(arr, root.Left, 2*i+1, n)
			root.Right = InsertLevelOrder(arr, root.Right, 2*i+2, n)
		}
	}
	return root
}

func Inorder(t *Node) []int {
	if t == nil {
		return []int{}
	}
	arr := Inorder(t.Left)
	arr = append(arr, t.Val)
	arr = append(arr, Inorder(t.Right)...)
	return arr
}

func Postorder(t *Node) []int {
	if t == nil {
		return []int{}
	}
	arr := Postorder(t.Left)
	arr = append(arr, Postorder(t.Right)...)
	arr = append(arr, t.Val)
	return arr
}
