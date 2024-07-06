package main

import (
	"fmt"
)

func main() {
	fmt.Println(distinctNumbers([]int{1, 2, 3, 2, 2, 1, 3}, 3))
}

func distinctNumbers(nums []int, k int) []int {
	mp := make(map[int]int)
	var (
		start  int
		result []int
	)

	for i := 0; i < len(nums); i++ {
		if i >= k {
			result = append(result, len(mp))
			mp[nums[start]]--
			if mp[nums[start]] == 0 {
				delete(mp, nums[start])
			}
			start++
		}

		mp[nums[i]]++
	}

	result = append(result, len(mp))

	return result
}
