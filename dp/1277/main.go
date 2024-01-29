package main

import "fmt"

func main() {
	fmt.Println(countSquares([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1}}))
}

func countSquares(matrix [][]int) int {
	var result int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				result++
			}
			result += checker(matrix, i, j)
			// fmt.Println()
		}
	}
	return result
}

func checker(matrix [][]int, x, y int) int {
	var count int
	xx, yy := x, y
	for x < len(matrix) && y < len(matrix[0]) {
		x++
		y++
		if x == len(matrix) || y == len(matrix[0]) {
			break
		}

		var check bool
		for i := xx; i <= x; i++ {
			for j := yy; j <= y; j++ {
				// fmt.Println(i, j)
				if matrix[i][j] != 1 {
					check = true
				}
				if check {
					break
				}
			}
			if check {
				break
			}
		}
		// fmt.Println(check, x, y, xx, yy)
		if !check {
			count++
		} else {
			break
		}

	}

	return count
}
