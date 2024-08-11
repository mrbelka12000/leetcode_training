package main

import (
	"fmt"
)

func main() {
	fmt.Println(countOfPairs([]int{5, 5, 5, 5}))
}

func countOfPairs(nums []int) int {
	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, 1001)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	return runner(nums, 0, 0, 0, 0, memo)
}

const mod = 1000000007

func runner(nums []int, a1, a2 int, pos, v2 int, memo [][]int) int {
	if pos == len(nums) {
		return 1
	}

	if val := memo[pos][v2]; val != -1 {
		return val
	}

	var count int

	for i := 0; i <= nums[pos]; i++ {
		if pos > 0 {
			if a1 > i || a2 < nums[pos]-i {
				continue
			}
			if a1+a2 != nums[pos-1] {
				break
			}
		}

		count = count + runner(nums, i, nums[pos]-i, pos+1, i, memo)%mod
	}

	count = count % mod

	memo[pos][v2] = count

	return count
}
