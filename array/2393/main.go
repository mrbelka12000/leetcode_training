package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubarrays([]int{1, 2, 3, 4, 5}))
}

func countSubarrays(nums []int) int64 {
	var (
		result, start, last int
	)

	last = -1
	for i := 0; i < len(nums); i++ {
		if last != -1 && last >= nums[i] {
			start = i
		}
		last = nums[i]

		result += i - start + 1
	}

	return int64(result)
}
