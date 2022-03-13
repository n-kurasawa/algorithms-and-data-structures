package main

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
		w = 15
	)
	values := []int{}
	wights := []int{}

	dp := make([][]int, n+1)

	for i := 0; i < n; i++ {
		for j := 0; j <= w; w++ {
			if j-wights[i] >= 0 {

			}
		}
	}

}
