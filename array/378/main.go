package main

import (
	"sort"
)

func main() {
}

func kthSmallest(matrix [][]int, k int) int {
	var result []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result = append(result, matrix[i][j])
		}
	}

	sort.Ints(result)

	// fmt.Println(result)

	return result[k-1]
}
