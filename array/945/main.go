package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	var result int

	for len(nums) != 1 {
		if nums[0] == nums[1] {
			prev := nums[0]
			for i := 1; i < len(nums); i++ {
				if nums[i] != prev {
					break
				}
				result++
				nums[i]++
			}
		}
		nums = nums[1:]
	}

	return result
}
