package relative_sorting

//package main

import (
	"fmt"
	"sort"
)

func solve(a []int, b []int) []int {
	ans := []int{}
	mp := map[int]int{}
	for i := 0; i < len(a); i++ {
		mp[a[i]]++
	}

	for i := 0; i < len(b); i++ {
		ans = append(ans, b[i])
		if _, ok := mp[b[i]]; ok {
			for j := 1; j < mp[b[i]]; j++ {
				ans = append(ans, b[i])
			}
			delete(mp, b[i])
		}
	}
	rest := []int{}
	for i := 0; i < len(a); i++ {
		if _, ok := mp[a[i]]; ok {
			rest = append(rest, a[i])
		}
	}
	sort.Ints(rest)
	for i := 0; i < len(rest); i++ {
		ans = append(ans, rest[i])
	}

	return ans
}

func main() {
	a := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8}
	b := []int{2, 1, 8, 3}
	fmt.Printf("%+v\n", solve(a, b)) // output for debug

}
