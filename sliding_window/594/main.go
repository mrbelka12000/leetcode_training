package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
	sort.Ints(nums)
	var result, start int
	for i := 1; i < len(nums); i++ {
		for nums[i]-nums[start] > 1 {
			start++
		}
		if nums[i]-nums[start] == 1 {
			result = max(result, i-start+1)
		}
	}

	return result
}
