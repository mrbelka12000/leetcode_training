package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}

func minFallingPathSum(matrix [][]int) int {
	mp := initMap(matrix)
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			minVal := math.MaxInt32
			for k := 0; k < len(matrix); k++ {
				if j == k {
					continue
				}
				jj := k
				val, ok := mp[i-1][jj]
				if !ok {
					continue
				}
				minVal = min(minVal, val+matrix[i][j])
			}
			mp[i][j] = minVal
		}
	}

	var result int = math.MaxInt32
	for _, v := range mp[len(matrix)-1] {
		result = min(result, v)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initMap(matrix [][]int) map[int]map[int]int {
	mp := make(map[int]map[int]int)

	for i := 0; i < len(matrix); i++ {
		mp[i] = make(map[int]int)
		mp[0][i] = matrix[0][i]
	}
	return mp
}
