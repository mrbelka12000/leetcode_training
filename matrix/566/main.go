package main

import (
	"fmt"
)

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat[0])*len(mat) != r*c {
		return mat
	}
	newMatrix := make([][]int, r)
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, c)
	}

	var x, y int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			newMatrix[x][y] = mat[i][j]
			y++
			if y == c {
				y = 0
				x++
			}
		}
	}

	return newMatrix
}
