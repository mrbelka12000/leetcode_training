package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func minPathSum(grid [][]int) int {
	memo := map[[2]int]int{
		[2]int{0, 0}: grid[0][0],
	}
	for i := 1; i < len(grid); i++ {
		memo[[2]int{i, 0}] = grid[i][0] + memo[[2]int{i - 1, 0}]
	}

	for j := 1; j < len(grid[0]); j++ {
		memo[[2]int{0, j}] = grid[0][j] + memo[[2]int{0, j - 1}]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			val := grid[i][j]
			// fmt.Println(val)
			memo[[2]int{i, j}] = min(val+memo[[2]int{i - 1, j}], val+memo[[2]int{i, j - 1}])
		}
	}

	// fmt.Println(memo)
	return memo[[2]int{len(grid) - 1, len(grid[0]) - 1}]
}

// [[1,2,3],
// [4,5,6]]
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
