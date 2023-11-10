package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(y int, x int) int {
	memo := make(map[[2]int]int)
	for i := 0; i < x; i++ {
		memo[[2]int{i, 0}] = 1
	}
	for i := 0; i < y; i++ {
		memo[[2]int{0, i}] = 1
	}

	for i := 1; i < x+1; i++ {
		for j := 1; j < y+1; j++ {
			memo[[2]int{i, j}] = memo[[2]int{i - 1, j}] + memo[[2]int{i, j - 1}]
		}
	}
	// fmt.Println(memo)
	return memo[[2]int{x - 1, y - 1}]
}
