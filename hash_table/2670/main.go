package main

import (
	"fmt"
)

func main() {
	fmt.Println(distinctDifferenceArray([]int{1,2,3,4,5}))
}

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suf := make([]int, n)

	mp := make(map[int]int)

	for i, v := range nums{
		mp[v]++
		suf[i] = len(mp)
	}


	result := make([]int, n)
	check := make(map[int]bool)
	for i := n-1 ;i >=0 ;i--{
		result[i] = len(mp) - len(check)
		mp[nums[i]]--
		if mp[nums[i]] == 0{
			delete(mp, nums[i])
		}
		check[nums[i]] = true
	}
	return result
}
}