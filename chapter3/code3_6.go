package main

import "fmt"

func exec(arr []int, w int) bool {
	for bit := 0; bit < 1<<len(arr); bit++ {
		var sum int
		for i := 0; i < len(arr); i++ {
			if (bit & (1 << i)) > 0 {
				sum += arr[i]
			}
		}
		if sum == w {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(exec([]int{1, 2, 4, 5, 11}, 10))
}
