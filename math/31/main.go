package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4}
	nextPermutation(arr)
	fmt.Println(arr)
}

func nextPermutation(nums []int) {
	k := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}

	// last permutation
	if k == -1 {
		sort.Ints(nums)
		return
	}

	var maxInd int
	for i := 0; i < len(nums); i++ {
		if nums[k] < nums[i] {
			maxInd = i
		}
	}

	nums[maxInd], nums[k] = nums[k], nums[maxInd]
	l, r := k+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
