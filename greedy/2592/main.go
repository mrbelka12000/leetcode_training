package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1}))
}

func maximizeGreatness(nums []int) int {
	cp := make([]int, len(nums))

	copy(cp, nums)
	sort.Ints(cp)
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	sort.Ints(nums)

	var result int
	for i := 0; i < len(nums); i++ {
		if len(cp) == 0 {
			break
		}

		for cp[0] <= nums[i] {
			cp = cp[1:]
			if len(cp) == 0 {
				return result
			}
		}

		if cp[0] > nums[i] {
			cp = cp[1:]
			result++
		}
	}

	return result
}
