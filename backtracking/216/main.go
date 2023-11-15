package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
}

func combinationSum3(k int, n int) [][]int {
	var result [][]int

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	solver(nums, []int{}, n, k, &result)

	return result
}

func solver(nums, curr []int, target, k int, result *[][]int) {
	if len(curr) > k {
		return
	}

	if target == 0 && len(curr) == k {
		temp := []int{}
		temp = append(temp, curr...)
		*result = append(*result, temp)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}

		solver(nums[i+1:], append(curr, nums[i]), target-nums[i], k, result)
	}
}
