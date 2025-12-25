package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	psMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		psMatrix[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			psMatrix[i][j] = matrix[i][j]
			if i > 0 {
				psMatrix[i][j] += psMatrix[i-1][j]
			}
			if j > 0 {
				psMatrix[i][j] += psMatrix[i][j-1]
			}
			if i > 0 && j > 0 {
				psMatrix[i][j] -= psMatrix[i-1][j-1]
			}
		}
	}

	// for _, v := range psMatrix{
	//     fmt.Println(v)
	// }

	return NumMatrix{
		matrix: psMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := this.matrix[row2][col2]

	if row1 != 0 {
		result -= this.matrix[row1-1][col2]
	}

	if col1 != 0 {
		result -= this.matrix[row2][col1-1]
	}

	if row1 != 0 && col1 != 0 {
		result += this.matrix[row1-1][col1-1]
	}

	return result
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
