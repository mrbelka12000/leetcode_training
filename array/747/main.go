package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}

func dominantIndex(nums []int) int {
	var maxInd, tmp int
	for i, v := range nums {
		if tmp < v {
			tmp = v
			maxInd = i
		}
	}

	sort.Ints(nums)
	if nums[len(nums)-1] >= nums[len(nums)-2]*2 {
		return maxInd
	}

	return -1
}
