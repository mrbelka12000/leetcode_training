package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
}

func numRescueBoats(nums []int, limit int) int {
	sort.Ints(nums)
	var result int
	for len(nums) > 1 {
		result++
		size := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		if size >= limit {
			continue
		}

		size += nums[0]
		if size > limit {
			continue
		}

		nums = nums[1:]
	}

	if len(nums) != 0 {
		result++
	}

	return result
}
