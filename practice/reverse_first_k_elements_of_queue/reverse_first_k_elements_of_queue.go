package reverse_first_k_elements_of_queue

// package REPLACE

import (
	"fmt"
	//	"github.com/nak3/note/lib"
)

// solve ...
func solve(n, k int, arr []int) []int {
	stk := []int{}
	for i := 0; i < k; i++ {
		stk = append(stk, arr[i])
	}
	for i := 0; i < k; i++ {
		arr[i] = stk[k-i-1]
	}
	return arr

}

func main() {
	n := 5
	k := 3
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%+v\n", solve(n, k, arr)) // output for debug

}
