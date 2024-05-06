package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathCost([][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}))
}

func minPathCost(grid [][]int, moveCost [][]int) int {
	tmp := grid[0]

	result := math.MaxInt32
	for i := 1; i < len(grid); i++ {
		newArr := getArr(len(grid[i]))

		for ind, v := range tmp {
			cost := moveCost[grid[i-1][ind]]
			for j := 0; j < len(grid[i]); j++ {
				newArr[j] = min(newArr[j], v+grid[i][j]+cost[j])
				if i == len(grid)-1 {
					result = min(result, newArr[j])
				}
			}
		}
		tmp = newArr
	}

	return result
}

func getArr(n int) []int {
	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		arr[i] = math.MaxInt32
	}

	return arr
}
