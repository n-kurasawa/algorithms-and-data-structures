package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	var (
		n = 6
		w = 9
	)
	wights := []int{2, 1, 3, 2, 1, 5}
	values := []int{3, 2, 6, 1, 3, 85}

	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if j >= wights[i] {
				dp[i+1][j] = max(dp[i][j], dp[i][j-wights[i]]+values[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	fmt.Println(dp[n][w])
}
