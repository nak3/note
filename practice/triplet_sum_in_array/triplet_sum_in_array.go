package triplet_sum_in_array

import (
	"fmt"
	"sort"
)

func solve(data []int, target int) int {
	sort.Ints(data)

	for i := 0; i < len(data); i++ {
		targ := target - data[i]
		l := i + 1
		r := len(data) - 1
		for l < r {
			tmp := data[l] + data[r]
			if tmp == targ {
				return 1
			}
			if tmp < l+r {
				l++
			}
			if tmp > l+r {
				r--
			}
		}
	}

	return 0
}

func main() {
	target := 13
	data := []int{1, 4, 45, 6, 10, 8}
	fmt.Printf("%+v\n", solve(data, target)) // output for debug
}
