package main

import "fmt"

func fibo(n int, memo []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fibo(n-1, memo) + fibo(n-2, memo)
	return memo[n]
}

func main() {
	memo := make([]int, 7)
	fmt.Println(fibo(6, memo))
}
