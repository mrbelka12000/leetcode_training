package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func rangeAddQueries(n int, queries [][]int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	for _, v := range queries {
		for i := v[0]; i <= v[2]; i++ {
			for j := v[1]; j <= v[3]; j++ {
				matrix[i][j]++
			}
		}
	}

	return matrix
}
