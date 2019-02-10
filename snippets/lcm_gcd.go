package main

import (
	"fmt"
)

// gcd ...
func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

// lcm ...
func lcm(a, b int) int {
	g := gcd(a, b)
	return (a / g) * b
}

// Verified by https://onlinejudge.u-aizu.ac.jp/courses/library/6/NTL/1/NTL_1_C
func main() {
	fmt.Printf("%v\n", gcd(12, 8))
	fmt.Printf("%v\n", lcm(4, 3))
}
