package kmp

//package main

import "fmt"

func KMPSearch(b, pattern string) int {
	ans := -1
	table := createTable(pattern)

	p := 0
	i := 0
	fmt.Printf("%+v\n", table) // output for debug

	for i < len(b) && p < len(pattern) {
		if b[i] == pattern[p] {
			i++
			p++
		} else if p == 0 {
			i++
		} else {
			p = table[p]
		}
	}
	if p == len(pattern) {
		return i - p
	}
	return ans
}

func createTable(pattern string) []int {
	table := make([]int, len(pattern))

	j := 0
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[j] {
			table[i] = j
			j++
		} else {
			table[i] = j
			j = 0
		}
	}
	return table
}

func main() {
	base := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"
	ans := KMPSearch(base, pattern)
	fmt.Printf("%+v\n", ans) // output for debug
}
