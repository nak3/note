package practice

import (
	//	"../../lib"

	"github.com/nak3/note/lib"
)

func constructMaximumBinaryTree(nums []int) *lib.Node {
	var left *lib.Node
	var right *lib.Node
	mx := 0
	var idx int
	for k, v := range nums {
		if v > mx {
			idx = k
			mx = v
		}
	}

	if idx == 0 {
		left = nil
	} else {
		left = constructMaximumBinaryTree(nums[:idx])
	}
	if idx == len(nums)-1 {
		right = nil
	} else {
		right = constructMaximumBinaryTree(nums[idx+1:])

	}
	return &lib.Node{mx, left, right}
}
