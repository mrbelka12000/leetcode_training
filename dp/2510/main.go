package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func isThereAPath(grid [][]int) bool {
	visited = make(map[[3]int]bool)
	return dfs(grid, 0, 0, 0, 0)
}

var visited map[[3]int]bool

func dfs(grid [][]int, z, o, x, y int) bool {
	if x >= len(grid) || y >= len(grid[x]) {
		return false
	}

	if grid[x][y] == 0 {
		z++
	} else {
		o++
	}

	if val, ok := visited[[3]int{x, y, z}]; ok {
		return val
	}

	if x == len(grid)-1 && y == len(grid[x])-1 {
		return z == o
	}

	visited[[3]int{x, y, z}] = dfs(grid, z, o, x, y+1) || dfs(grid, z, o, x+1, y)

	return visited[[3]int{x, y, z}]
}

/*
[[1,1,1,1,1,1,1,1,1]
,[1,1,0,0,1,0,0,0,1]]
*/
