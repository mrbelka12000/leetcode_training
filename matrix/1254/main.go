package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func closedIsland(grid [][]int) int {
	var result int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				check = false
				runner(grid, i, j)
				if !check {
					result++
				}
			}
		}
	}

	return result
}

var check bool

func runner(grid [][]int, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
		check = true
		return
	}

	if grid[x][y] == 1 {
		return
	}

	grid[x][y] = 1

	runner(grid, x+1, y)
	runner(grid, x-1, y)
	runner(grid, x, y+1)
	runner(grid, x, y-1)
}
