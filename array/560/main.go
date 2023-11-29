package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func subarraySum(nums []int, k int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		if tmp == k {
			result++
		}
		for j := i + 1; j < len(nums); j++ {
			tmp += nums[j]
			if tmp == k {
				result++
			}
			// fmt.Println(tmp,nums[i], nums[j])
		}
	}

	return result
}
