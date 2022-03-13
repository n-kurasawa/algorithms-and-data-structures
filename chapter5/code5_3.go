package main

import (
	"fmt"
	"math"
)

func exec(h []int) int {
	dp := make([]int, len(h))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < len(h); i++ {
		chmin(dp, i, dp[i-1]+abs(h[i]-h[i-1]))
		if i > 1 {
			chmin(dp, i, dp[i-2]+abs(h[i]-h[i-2]))
		}
	}
	return dp[len(h)-1]
}

func chmin(dp []int, i, b int) {
	if dp[i] > b {
		dp[i] = b
	}
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

func main() {
	fmt.Println(exec([]int{2, 9, 4, 5, 1, 6, 10}))
}
