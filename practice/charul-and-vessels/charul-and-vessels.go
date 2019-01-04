package charul_and_vessels

import (
	"fmt"
	"sort"
)

func dp(target int, k []int) bool {
	sort.Ints(k)
	fmt.Printf("%+v\n", k) // output for debug
	mp := map[int]struct{}{}
	mp[0] = struct{}{}

	for i := 1; i <= target; i++ {
		for _, v := range k {
			if _, ok := mp[i-v]; ok {
				mp[i] = struct{}{}
			}
		}
	}
	_, ok := mp[target]
	return ok
}

func main() {
	k := []int{6, 3, 4, 2, 1}
	target := 20
	ans := dp(target, k)
	fmt.Printf("%+v\n", ans) // output for debug
}
