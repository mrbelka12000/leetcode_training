package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func numMagicSquaresInside(grid [][]int) int {
	var result int

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid)-2; j++ {
			if runner(grid, i, j) {
				result++
			}
		}
	}

	return result
}

func runner(grid [][]int, x, y int) bool {
	mp := make(map[int]bool)
	for i := x; i < x+3 && i < len(grid); i++ {
		for j := y; j < y+3 && j < len(grid[i]); j++ {
			mp[grid[i][j]] = true
		}
	}

	for i := 1; i <= 9; i++ {
		if !mp[i] {
			return false
		}
	}

	var rowSum, colSum int
	for i := x; i < x+3 && i < len(grid); i++ {
		var tmp int
		for j := y; j < y+3 && j < len(grid[i]); j++ {
			tmp += grid[i][j]
		}
		if rowSum == 0 {
			rowSum = tmp
		}
		if rowSum != tmp {
			return false
		}
	}

	for j := y; j < y+3 && j < len(grid[0]); j++ {
		var tmp int
		for i := x; i < x+3 && i < len(grid); i++ {
			tmp += grid[i][j]
		}

		if colSum == 0 {
			colSum = tmp
		}
		if colSum != tmp {
			return false
		}
	}

	return colSum == rowSum
}
