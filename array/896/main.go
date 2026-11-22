package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMonotonic([]int{1, 3, 2, 4}))
}

func isMonotonic(nums []int) bool {
	return check(nums, true) || check(nums, false)
}

func check(arr []int, f bool) bool {
	for i := 1; i < len(arr); i++ {
		if f {
			if arr[i] > arr[i-1] {
				return false
			}
		} else {
			if arr[i] < arr[i-1] {
				return false
			}
		}
	}
	return true
}
