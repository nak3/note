package array_subset_of_another_array

//package main

import (
	"fmt"
	"sort"
)

func solve(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	idx := 0
	ok := true
	for i := 0; i < len(b); i++ {
		target := b[i]
		l := idx
		r := len(a) - 1
		ok = false
		for l <= r {
			mid := (l + r) >> 1
			if a[mid] == target {
				idx = mid
				ok = true
				break
			}
			if l == r {
				break
			}
			if a[mid] > target {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if ok == false {
			break
		}
		idx++
	}

	return ok

}

func main() {
	a := []int{11, 1, 13, 21, 3, 7}
	b := []int{11, 3, 7, 13}
	fmt.Printf("%+v\n", solve(a, b)) // output for debug

}
