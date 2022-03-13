package main

import "fmt"

func exec(i, w int, arr []int) bool {
	if i == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	if exec(i-1, w, arr) {
		return true
	}

	if exec(i-1, w-arr[i-1], arr) {
		return true
	}
	return false
}

func main() {
	fmt.Println(exec(4, 14, []int{3, 2, 6, 5}))
}
