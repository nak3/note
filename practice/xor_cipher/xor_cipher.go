package xor_cipher

//package main

import (
	"fmt"
	"strconv"
)

func solve(d string) string {
	ans := ""
	xor := 0

	l := len(d) / 2
	if len(d)%2 != 0 {
		l += 1
	}

	prev := 0
	for i := 0; i < l; i++ {
		x, _ := strconv.ParseInt(string(d[i]), 16, 64)
		xor = int(x) ^ prev
		prev ^= xor
		ans += fmt.Sprintf("%x", xor)
	}

	return ans
}

func main() {
	//d := "A1D0A1D"
	d := "653CAE8DA8EDB426052"
	fmt.Printf("%+v\n", solve(d)) // output for debug
}
