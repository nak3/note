package paths_to_reach_origin

// package REPLACE

import (
	"fmt"
	//	"github.com/nak3/note/lib"
)

// solve ...
func solve(n, m int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n][m]
}

func main() {
	n := 3
	m := 2
	fmt.Printf("%+v\n", solve(n, m)) // output for debug
}
