package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))

}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	mp := make(map[[4]int]struct{})

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			n1, n2 := nums[i], nums[j]
			n3, n4 := j+1, len(nums)-1

			for n3 < n4 {
				if _, ok := mp[[4]int{n1, n2, nums[n3], nums[n4]}]; ok {
					n3++
					continue
				}

				mid := n1 + n2 + nums[n3] + nums[n4]
				if mid < target {
					n3++
				} else if mid > target {
					n4--
				} else {
					result = append(result, []int{n1, n2, nums[n3], nums[n4]})
					mp[[4]int{n1, n2, nums[n3], nums[n4]}] = struct{}{}
				}
			}
		}
	}

	return result
}
