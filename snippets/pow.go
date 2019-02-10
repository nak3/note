package main

import (
	"fmt"
)

func pow(m, n, mod int) int {
	ans := 1
	if n > 0 {
		if n%2 == 0 {
			ans = pow(m*m%mod, n/2, mod)
		} else {
			ans = pow(m*m%mod, n/2, mod) * m % mod
		}
	}
	return ans
}

// Verified by https://onlinejudge.u-aizu.ac.jp/courses/library/6/NTL/1/NTL_1_B
/*
func main() {
	m, n, mod := 2, 10, 1000000007
	fmt.Printf("%+v\n", pow(m, n, mod)) // output for debug
}
*/
