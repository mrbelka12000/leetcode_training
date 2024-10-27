package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumOperations([]int{4, 3, 2, 1, 2, 3, 1}))
}

func minimumOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var result int
	for len(nums) > 1 {
		l, r := 0, len(nums)-1
		if nums[l] == nums[r] {
			nums = nums[1:]
			nums = nums[:len(nums)-1]
			continue
		}

		if nums[l] < nums[r] {
			nums[1] += nums[l]
			nums = nums[1:]
			result++
		} else {
			nums[r-1] += nums[r]
			nums = nums[:r]
			result++
		}
	}

	return result
}
