package main

import "fmt"

func main() {
	fmt.Println(uniquePathsIII([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}))
}

func uniquePathsIII(grid [][]int) int {
	empty = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				empty++
			}
		}
	}

	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				solver(grid, 0, i, j, &result)
			}
		}
	}

	return result
}

var dirs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
var empty int

func solver(grid [][]int, count, x, y int, result *int) {
	// fmt.Println(x,y, count)
	if grid[x][y] == 2 {
		if count == empty+1 {
			*result++
		}
		return
	}

	grid[x][y] = 1

	// fmt.Println(grid)

	for _, v := range dirs {
		xx, yy := x+v[0], y+v[1]
		if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[xx]) || grid[xx][yy] == -1 || grid[xx][yy] == 1 {
			continue
		}

		solver(grid, count+1, xx, yy, result)

		if grid[xx][yy] != 2 {
			grid[xx][yy] = 0
		}
	}
}
