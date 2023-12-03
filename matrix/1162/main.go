package main

import "fmt"

func main() {

}

func maxDistance(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = -1
			}
		}
	}

	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == -1 {
				mark(grid, i, j, i, j, make(map[[2]int]struct{}))
				// print(grid)
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > result {
				result = grid[i][j]
			}
		}
	}

	if result == 0 {
		return -1
	}

	return result
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func mark(grid [][]int, orgX, orgY, x, y int, mp map[[2]int]struct{}) {

	mp[[2]int{x, y}] = struct{}{}
	for _, dir := range dirs {

		xx, yy := x+dir[0], y+dir[1]
		if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[0]) || grid[xx][yy] == -1 {
			continue
		}

		if _, ok := mp[[2]int{xx, yy}]; ok {
			continue
		}

		grid[xx][yy] = min(grid[xx][yy], getDist(orgX, orgY, xx, yy))

		mark(grid, orgX, orgY, xx, yy, mp)
	}
}

func getDist(x, y, x1, y1 int) int {
	return abs(x-x1) + abs(y-y1)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func print(grid [][]int) {
	fmt.Println()
	for _, v := range grid {
		fmt.Println(v)
	}

	fmt.Println()
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if a < b {
		return a
	}
	return b
}
