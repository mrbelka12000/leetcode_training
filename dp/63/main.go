package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func uniquePathsWithObstacles(grid [][]int) int {
	y, x := len(grid[0])-1, len(grid)-1

	memo := make(map[[2]int]int)

	for i := 0; i < x+1; i++ {
		if grid[i][0] == 1 {
			break
		}
		memo[[2]int{i, 0}] = 1
	}

	for i := 0; i < y+1; i++ {
		if grid[0][i] == 1 {
			break
		}
		memo[[2]int{0, i}] = 1
	}

	for i := 1; i < x+1; i++ {
		for j := 1; j < y+1; j++ {
			if grid[i][j] == 1 {
				continue
			}
			var sum int
			if grid[i-1][j] != 1 {
				sum += memo[[2]int{i - 1, j}]
			}
			if grid[i][j-1] != 1 {
				sum += memo[[2]int{i, j - 1}]
			}
			// fmt.Println(i, j, sum)
			memo[[2]int{i, j}] = sum
		}
	}

	// fmt.Println(memo)

	return memo[[2]int{x, y}]
}
