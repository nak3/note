package equalize_the_towers

//package main

import (
	"fmt"
	"math"
)

func costCalc(h []int, cost []int, mid int) int {
	sum := 0
	for i := 0; i < len(h); i++ {
		sum += int(math.Abs(float64(h[i]-mid))) * cost[i]

	}
	return sum
}

func solve(h []int, cost []int) int {
	r, l := 0, 0
	ans := 100001
	for i := 0; i < len(h); i++ {
		if h[i] > r {
			r = h[i] + 1
		}
	}
	l = 0
	for r > l {
		mid := (l + r) / 2
		var bm int

		if mid > 0 {
			bm = costCalc(h, cost, mid-1)
		} else {
			bm = 100001
		}
		m := costCalc(h, cost, mid)
		am := costCalc(h, cost, mid+1)
		if ans == m {
			break
		}
		if ans > m {
			ans = m
		}
		if bm <= m {
			r = mid
		} else if am <= m {
			l = mid + 1
		}
	}
	return ans
}

func main() {
	h := []int{1, 2, 3}
	c := []int{10, 100, 1000}
	fmt.Printf("%+v\n", solve(h, c)) // output for debug
}
