package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestMagicSquare([][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}))
}

func largestMagicSquare(grid [][]int) int {
	R, C := len(grid), len(grid[0])

	row := make([][]int, R)
	col := make([][]int, R+1)
	for i := 0; i < R; i++ {
		row[i] = make([]int, C+1)
	}
	for i := 0; i <= R; i++ {
		col[i] = make([]int, C)
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			row[i][j+1] = row[i][j] + grid[i][j]
			col[i+1][j] = col[i][j] + grid[i][j]
		}
	}

	ans := 1

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			maxSize := min(R-i, C-j)
			for size := maxSize; size > ans; size-- {
				if isMagic(grid, row, col, i, j, size) {
					ans = size
					break
				}
			}
		}
	}
	return ans
}

func isMagic(g, r, c [][]int, x, y, l int) bool {
	target := r[x][y+l] - r[x][y]

	for i := 0; i < l; i++ {
		if r[x+i][y+l]-r[x+i][y] != target {
			return false
		}
	}

	for j := 0; j < l; j++ {
		if c[x+l][y+j]-c[x][y+j] != target {
			return false
		}
	}

	d1, d2 := 0, 0
	for k := 0; k < l; k++ {
		d1 += g[x+k][y+k]
		d2 += g[x+l-1-k][y+k]
	}
	return d1 == target && d2 == target
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
