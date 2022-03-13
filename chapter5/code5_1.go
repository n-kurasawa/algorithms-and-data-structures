package main

import "fmt"

func exec(h []int) int {
	dp := make([]int, len(h))
	dp[0] = 0
	for i := 1; i < len(h); i++ {
		if i == 1 {
			dp[i] = abs(h[i] - h[i-1])
		} else {
			dp[i] = min(
				dp[i-1]+abs(h[i]-h[i-1]),
				dp[i-2]+abs(h[i]-h[i-2]),
			)
		}
	}
	return dp[len(h)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
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
