package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var result [][]int
	runner(nums, nil, k, &result)
	return result
}

func runner(nums, tmp []int, k int, result *[][]int) {
	fmt.Println(nums, tmp)
	if len(tmp) == k {
		check := make([]int, k)
		copy(check, tmp)
		*result = append(*result, check)
		return
	}

	for i := 0; i < len(nums); i++ {
		//newNums := make([]int, len(nums))
		//copy(newNums, nums)

		runner(nums[i+1:], append(tmp, nums[i]), k, result)
	}
}
