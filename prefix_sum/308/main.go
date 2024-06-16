package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type NumMatrix struct {
	prefSum [][]int
	matrix  [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefSum := make([][]int, len(matrix))
	for i := 0; i < len(prefSum); i++ {
		prefSum[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			prefSum[i][j] += matrix[i][j]
			if j != 0 {
				prefSum[i][j] += prefSum[i][j-1]
			}
		}
	}

	return NumMatrix{prefSum, matrix}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	d := val - this.matrix[row][col]
	for i := col; i < len(this.prefSum[0]); i++ {
		this.prefSum[row][i] += d
	}

	this.matrix[row][col] = val

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int

	for i := row1; i <= row2; i++ {
		sum += this.prefSum[i][col2]
		if col1 != 0 {
			sum -= this.prefSum[i][col1-1]
		}
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
