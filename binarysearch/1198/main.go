package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func smallestCommonElement(mat [][]int) int {
	if len(mat) == 1 {
		return mat[0][0]
	}

	for _, v := range mat[0] {
		count := 1

		for j := 1; j < len(mat); j++ {
			val := isExist(mat[j], v)
			if val {
				count++
			} else {
				break
			}
		}
		if count == len(mat) {
			return v
		}
	}

	return -1
}

func isExist(arr []int, target int) bool {
	l, r := -1, len(arr)-1
	for r-l > 1 {
		mid := (l + r) / 2
		if arr[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}
	return arr[r] == target
}
