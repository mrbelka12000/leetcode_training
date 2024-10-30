package main

import (
	"fmt"
)

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func transpose(matrix [][]int) [][]int {
	newMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			newMatrix[j][i] = matrix[i][j]
		}
	}

	return newMatrix
}
