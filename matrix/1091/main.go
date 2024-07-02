package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func shortestPathBinaryMatrix(grid [][]int) int {
	return bfs(grid)
}

func bfs(grid [][]int) int {
	dirs := [8][2]int{
		{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1},
	}

	q := [][2]int{{0, 0}}
	result := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			if n[0] < 0 || n[0] >= len(grid) || n[1] < 0 || n[1] >= len(grid[n[0]]) || grid[n[0]][n[1]] == 1 {
				continue
			}

			if n[0] == len(grid)-1 && n[1] == len(grid[0])-1 && grid[n[0]][n[1]] == 0 {
				return result
			}

			grid[n[0]][n[1]] = 1

			for _, dir := range dirs {
				q = append(q, [2]int{dir[0] + n[0], dir[1] + n[1]})
			}
		}
		result++
	}

	return -1
}
