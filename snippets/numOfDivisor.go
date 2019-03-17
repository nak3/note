package main

import (
//	"fmt"
// "strings"
)

// numDivisor returns number of divisors
func numOfDivisor(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		for j := i; j < n+1; j += i {
			arr[j] += 1
		}
	}
	return arr
}

// func main() {
// 	fmt.Printf("%+v\n", numOfDivisor(10)) // output for debug
// 	// [0 1 2 2 3 2 4 2 4 3 4]
// }
