package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func wallsAndGates(grid [][]int) {
	var q [][]int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}

	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	var dist int
	for len(q) > 0 {

		size := len(q)

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			x, y := n[0], n[1]

			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || visited[x][y] || grid[x][y] == -1 {
				continue
			}

			visited[x][y] = true

			grid[x][y] = dist

			q = append(q, []int{x + 1, y})
			q = append(q, []int{x - 1, y})
			q = append(q, []int{x, y + 1})
			q = append(q, []int{x, y - 1})
		}

		dist++
	}
}
