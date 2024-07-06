package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubarrays([]int{2, 1, 4, 3, 5}, 10))
}

func countSubarrays(nums []int, k int64) int64 {
	var (
		result int64
		start  int
		tmp    int
	)

	for i := 0; i < len(nums); i++ {
		tmp += nums[i]

		for int64(tmp*(i-start+1)) >= k {
			tmp -= nums[start]
			start++
		}

		result += int64(i - start + 1)
	}

	return result
}
