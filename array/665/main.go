package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPossibility([]int{5, 7, 1, 8}))
}

func checkPossibility(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			tmp := nums[i]
			nums[i] = nums[i+1]
			if runner(nums) {
				return true
			}
			nums[i] = tmp

			nums[i+1] = nums[i]
			if runner(nums) {
				return true
			}

			return false
		}
	}

	return true
}

func runner(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
