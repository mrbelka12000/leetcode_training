package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
}

func numSubarraysWithSum(nums []int, goal int) int {
	var sum, count int

	mp := map[int]int{
		0: 1,
	}

	for _, v := range nums {
		sum += v
		if val, ok := mp[sum-goal]; ok {
			count += val
		}
		mp[sum]++
	}

	return count
}
