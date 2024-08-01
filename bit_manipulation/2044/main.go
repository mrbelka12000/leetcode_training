package main

import (
	"fmt"
)

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func countMaxOrSubsets(nums []int) int {
	mp = make(map[int]int)
	maxB = 0

	for i, v := range nums {
		runner(nums, i+1, v)
	}

	return mp[maxB]
}

var (
	maxB int
	mp   map[int]int
)

func runner(nums []int, pos, curr int) {
	mp[curr]++
	if pos == len(nums) {
		maxB = max(maxB, curr)
		return
	}

	for i := pos; i < len(nums); i++ {
		runner(nums, i+1, curr|nums[i])
	}
}
