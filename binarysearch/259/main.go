package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3, -5}, 2))
}

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	var (
		result int
		n      = len(nums)
	)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			l, r := j+1, n-1
			for l <= r {
				mid := (l + r) / 2
				if nums[mid]+nums[i]+nums[j] < target {
					result++
					l++
				} else {
					r = mid - 1
				}
			}
		}
	}

	return result
}
