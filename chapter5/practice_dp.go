package main

import "fmt"

func main() {
	a := []int{7, -6, 9, -5, 10}
	n := len(a)
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		dp[i+1] = max(dp[i], dp[i]+a[i])
	}
	fmt.Println(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
