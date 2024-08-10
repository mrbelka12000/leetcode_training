package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if bSearch(matrix[i], target) {
			return true
		}
	}

	return false
}

func bSearch(arr []int, target int) bool {
	if arr[0] > target || arr[len(arr)-1] < target {
		return false
	}

	l, r := -1, len(arr)-1

	for r-l > 1 {
		mid := (r + l) / 2

		if arr[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}

	return arr[r] == target
}
