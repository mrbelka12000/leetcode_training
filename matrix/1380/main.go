package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func luckyNumbers(matrix [][]int) []int {
	var result []int

	var inds [][2]int

	for _, row := range matrix {
		var minVal, minInd int
		minVal = row[0]

		for i, col := range row {
			if col < minVal {
				minVal = col
				minInd = i
			}
		}

		inds = append(inds, [2]int{minVal, minInd})
	}

	for _, v := range inds {
		var check bool
		for i := 0; i < len(matrix); i++ {
			if matrix[i][v[1]] > v[0] {
				check = true
			}
		}

		if !check {
			result = append(result, v[0])
		}
	}

	return result
}
