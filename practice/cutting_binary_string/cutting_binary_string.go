package cutting_binary_string

//package main

import (
	"fmt"
	"math"
)

func isPowerOfFive(val int) bool {
	if val == 1 {
		return false
	}
	for val > 1 {
		if val%5 != 0 {
			return false
		}
		val /= 5
	}
	return true

}

func biToInt(data string) int {
	ret := 0
	j := 0.0
	for i := len(data) - 1; i >= 0; i-- {
		if string(data[i]) == "1" {
			ret += int(math.Pow(2.0, j))
		}
		j++
	}
	return ret
}

func solve(data string) int {
	dp := make([]int, len(data)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 99999
	}
	dp[len(data)-1] = -1

	for i := len(data) - 2; i >= 0; i-- {
		if string(data[i]) == "0" {
			continue
		}
		for j := i; j < len(data); j++ {
			if isPowerOfFive(biToInt(data[i : j+1])) {
				if j == len(data)-1 {
					dp[i] = 1
				}
				if dp[j+1] != 99999 {
					if dp[i] > dp[j+1]+1 {
						dp[i] = dp[j+1] + 1
					}
				}
			}
		}
	}
	return dp[0]
}

func main() {
	//	base := "101101101"
	base := "101101"
	fmt.Printf("%+v\n", solve(base)) // output for debug
}
