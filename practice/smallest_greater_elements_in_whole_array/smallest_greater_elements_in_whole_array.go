package smallest

//package main

import (
	"fmt"
	"sort"
	"strconv"
)

func solve(data []int) []string {
	original := make([]int, len(data))
	copy(original, data)
	sort.Ints(data)
	mp := map[int]int{}

	for i := 0; i < len(data); i++ {
		mp[data[i]] = i
	}

	ans := []string{}
	for i := 0; i < len(original); i++ {
		tmp := mp[original[i]] + 1
		s := "_"
		if tmp != len(original) {
			s = strconv.Itoa(data[tmp])
		}
		ans = append(ans, s)
	}

	return ans
}

func main() {
	data := []int{6, 3, 9, 8, 10, 2, 1, 15, 7}
	fmt.Printf("%+v\n", solve(data)) // output for debug
}
