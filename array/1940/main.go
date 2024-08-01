package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonSubsequence([][]int{{2, 3, 6, 8}, {1, 2, 3, 5, 6, 7, 10}, {2, 3, 4, 6, 9}}))
}

func longestCommonSubsequence(array [][]int) []int {
	var (
		result []int
		arr    [101]int
		maxLen int
	)

	for _, v := range array {
		maxLen = max(maxLen, len(v))
	}

	for i := 0; i < maxLen; i++ {
		for j := 0; j < len(array); j++ {
			if i >= len(array[j]) {
				continue
			}

			arr[array[j][i]]++

			if arr[array[j][i]] == len(array) {
				result = append(result, array[j][i])
			}
		}
	}

	return result
}
