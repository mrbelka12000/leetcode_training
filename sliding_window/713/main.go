package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	var result, tmp, start int
	tmp = 1

	for i := 0; i < len(nums); i++ {
		tmp *= nums[i]

		for tmp >= k && start <= i {
			tmp /= nums[start]
			start++
		}

		result += i - start + 1
	}

	return result
}
