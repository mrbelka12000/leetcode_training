package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))
}

func maximumScore(nums []int, multipliers []int) int {

	minVal := -1 * math.MaxInt32
	memo := make([][]int, len(multipliers))
	for i := range memo {
		memo[i] = make([]int, len(nums))
		for j := range memo[i] {
			memo[i][j] = minVal
		}
	}

	var runner func(i, l int) int
	runner = func(i, l int) int {
		if i == len(multipliers) {
			return 0
		}

		if val := memo[i][l]; val != minVal {
			return val
		}

		r := len(nums) - 1 - (i - l)
		m := multipliers[i]

		memo[i][l] = max(m*nums[l]+runner(i+1, l+1), m*nums[r]+runner(i+1, l))
		return memo[i][l]
	}

	return runner(0, 0)
}
