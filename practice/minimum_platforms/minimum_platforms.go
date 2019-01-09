package minimum_platforms

//package main

import (
	"fmt"
	"sort"
)

func solve(a, d []int) int {
	mp := map[int][]int{}
	for i := 0; i < len(a); i++ {
		mp[a[i]] = append(mp[a[i]], 1)
		mp[d[i]] = append(mp[d[i]], -1)
	}
	arr := []int{}
	for k, _ := range mp {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	cnt := 0
	mx := 0
	for i := 0; i < len(arr); i++ {
		tmp := mp[arr[i]]
		for j := 0; j < len(tmp); j++ {
			cnt += tmp[j]
			if cnt > mx {
				mx = cnt
			}
		}
	}
	return mx
}

func main() {
	a := []int{900, 940, 950, 1100, 1500, 1800}
	d := []int{910, 1200, 1120, 1130, 1900, 2000}
	fmt.Printf("%+v\n", solve(a, d))
}
