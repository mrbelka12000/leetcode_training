package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func countSubIslands(grid1 [][]int, grid2 [][]int) (result int) {
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[i]); j++ {
			if grid2[i][j] == 1 {
				check = true
				floodFill(grid1, grid2, i, j)
				if check {
					result++
				}
			}
		}
	}

	return result
}

var check bool

func floodFill(grid1 [][]int, grid2 [][]int, x, y int) {
	if x < 0 || x >= len(grid1) || y < 0 || y >= len(grid1[x]) || grid2[x][y] == 0 {
		return
	}

	if grid2[x][y] == 1 && grid1[x][y] == 0 {
		check = false
	}

	grid2[x][y] = 0

	floodFill(grid1, grid2, x+1, y)
	floodFill(grid1, grid2, x-1, y)
	floodFill(grid1, grid2, x, y+1)
	floodFill(grid1, grid2, x, y-1)
}
