package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))
}

func largestDivisibleSubset(nums []int) []int {
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(nums)+1)
	}

	result = nil
	sort.Ints(nums)
	runner(nums, nil, 0, dp)

	return result
}

var (
	result []int
)

func runner(nums, curr []int, pos int, dp [][]bool) {
	if len(result) < len(curr) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		result = tmp
	}

	if len(nums) == pos || dp[pos][len(curr)] {
		return
	}
	dp[pos][len(curr)] = true

	for i := pos; i < len(nums); i++ {
		if len(curr) == 0 {
			runner(nums, append(curr, nums[i]), i+1, dp)
		} else {
			v := curr[len(curr)-1]
			if v%nums[i] == 0 || nums[i]%v == 0 {
				runner(nums, append(curr, nums[i]), i+1, dp)
			}
		}
	}
}
