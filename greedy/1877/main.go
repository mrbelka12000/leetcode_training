package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minPairSum([]int{4, 1, 5, 1, 2, 5, 1, 5, 5, 4}))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)

	var result int
	l, r := 0, len(nums)-1
	for l < r {
		result = max(nums[l]+nums[r], result)
		l++
		r--
	}

	return result
}
