package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

func subsetXORSum(nums []int) int {
	var result int

	for i, v := range nums {
		result += summer(nums, i+1, v)
	}

	return result
}

func summer(nums []int, pos, curr int) (result int) {
	if pos == len(nums) {
		return curr
	}
	result += curr

	for i := pos; i < len(nums); i++ {
		result += summer(nums, i+1, curr^nums[i])
	}

	return result
}
