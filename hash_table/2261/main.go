package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countDistinct([]int{2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3}, 2, 2))
}

func countDistinct(nums []int, k int, p int) int {
	var (
		result int
		mp     = make(map[string]bool)
	)

	for i := 0; i < len(nums); i++ {
		var (
			count int
			key   strings.Builder
		)
		for j := i; j < len(nums); j++ {
			if nums[j]%p == 0 {
				count++
			}
			if count > k {
				break
			}

			key.WriteByte(byte(nums[j]) + '0')
			if !mp[key.String()] {
				mp[key.String()] = true
				result++
			}
		}
	}

	return result
}
