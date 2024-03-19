package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))

}

func findTargetSumWays(nums []int, target int) int {
	var result int

	runner(nums[1:], target-nums[0], &result)
	runner(nums[1:], target+nums[0], &result)

	return result
}

func runner(nums []int, target int, result *int) {
	if len(nums) == 0 {
		if target == 0 {
			*result++
		}
		return
	}

	runner(nums[1:], target-nums[0], result)
	runner(nums[1:], target+nums[0], result)
}
