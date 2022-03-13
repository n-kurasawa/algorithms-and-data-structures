package main

import "fmt"

func fibo(n int) int {
	f := make([]int, n)
	f[0] = 0
	f[1] = 1
	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n-1]
}

func main() {
	fmt.Println(fibo(7))
}
