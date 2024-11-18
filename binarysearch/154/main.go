package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{3, 3, 1, 2, 3, 5}))
}

func findMin(nums []int) int {
	return nums[findPeekInd(nums)]

	return 0
}

func findPeekInd(nums []int) int {
	l, r := 0, len(nums)-1
	var ind int
	for l < r {
		mid := (l + r) / 2
		if mid == 0 || mid == len(nums)-1 {
			return ind
		}

		if nums[mid] < nums[mid+1] {
			ind = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ind
}
