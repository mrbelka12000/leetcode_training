package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	ps := make([]int, n)
	ps[0] = nums[0]

	for i := 1; i < n; i++ {
		ps[i] = ps[i-1] + nums[i]
		if ps[i]%k == 0 {
			return true
		}
	}

	fmt.Println(ps)
	mp := make(map[int]int)
	for i, v := range ps {
		kk := v % k
		// fmt.Println(kk)
		if val, ok := mp[kk]; ok {
			if i-val == 1 {
				// fmt.Println("suka")
				continue
			}
			return true
		}
		mp[kk] = i
	}
	fmt.Println(mp)

	return false
}
