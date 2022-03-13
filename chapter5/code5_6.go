package main

import (
	"fmt"
	"math"
)

func exec(i int, h, dp []int) int {
	if dp[i] < math.MaxInt {
		return dp[i]
	}
	if i == 0 {
		return 0
	}
	chmin(dp, i, exec(i-1, h, dp)+abs(h[i]-h[i-1]))
	if i > 1 {
		chmin(dp, i, exec(i-2, h, dp)+abs(h[i]-h[i-2]))
	}
	return dp[i]
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
	h := []int{2, 9, 4, 5, 1, 6, 10}
	dp := make([]int, len(h))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	fmt.Println(exec(len(h)-1, h, dp))
}
