package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumRemoteness([][]int{{-1, 1, -1}, {5, -1, 4}, {-1, 3, -1}}))
}

func sumRemoteness(grid [][]int) int64 {
	var (
		cells = make(map[[2]int]int)
		sum   int
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != -1 {
				cells[[2]int{i, j}] = len(cells)
				sum += grid[i][j]
			}
		}
	}

	uf := make([]int, len(cells))
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}

	visited := make(map[[2]int]bool)
	sumArr := make([]int, len(cells))
	for k, v := range cells {
		connect(grid, cells, visited, uf, sumArr, k[0], k[1], v)
	}

	var result int64

	for i := range sumArr {
		result += int64(sum - sumArr[find(uf, i)])
	}

	return result
}

var dirs = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func connect(grid [][]int, cells map[[2]int]int, visited map[[2]int]bool, uf, sum []int, x, y, idx int) {
	xf := find(uf, idx)
	if !visited[[2]int{x, y}] {
		sum[xf] += grid[x][y]
		visited[[2]int{x, y}] = true
	}

	for _, v := range dirs {
		xx, yy := x+v[0], y+v[1]

		if !visited[[2]int{xx, yy}] {
			n, ok := cells[[2]int{xx, yy}]
			if !ok {
				continue
			}
			sum[xf] += grid[xx][yy]

			visited[[2]int{xx, yy}] = true

			union(uf, xf, find(uf, n))
			connect(grid, cells, visited, uf, sum, xx, yy, n)
		}
	}
}

func union(uf []int, x, y int) {
	uf[y] = x
}

func find(uf []int, x int) int {
	for uf[x] != x {
		x = uf[x]
	}

	return x
}
