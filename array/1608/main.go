package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func specialArray(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	l, r := -1, nums[n-1]
	// fmt.Println(nums)

	l = -1

	for i := 0; i <= r; i++ {
		var gt int
		for ind, val := range nums {
			if val >= i {
				if n-ind+gt == i {
					l = i
				}
				gt++
				break
			}
		}
	}

	return l
}
