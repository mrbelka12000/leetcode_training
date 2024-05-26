package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v]++
	}

	var nums [][2]int

	for k, v := range mp {
		nums = append(nums, [2]int{k, v})
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})

	for k > 0 {
		if nums[0][1] <= k {
			k -= nums[0][1]
			nums = nums[1:]
			continue
		}
		break
	}

	return len(nums)
}
