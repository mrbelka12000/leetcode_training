package main

import (
	"fmt"
)

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
}

func zeroFilledSubarray(nums []int) int64 {
	var (
		result int
		start  int = -1
		found  bool
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if !found {
				start = i
			}
			found = true
			result += i - start + 1
		} else {
			found = false
		}
	}

	return int64(result)
}
