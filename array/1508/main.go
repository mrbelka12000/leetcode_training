package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, len(nums)*len(nums))

	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		sums = append(sums, tmp)
		for j := i + 1; j < len(nums); j++ {
			tmp += nums[j]
			sums = append(sums, tmp)
		}
	}

	sort.Ints(sums)

	var result int

	for i := left - 1; i <= right-1; i++ {
		result = (result + sums[i]) % mod
	}

	return result % mod
}

const mod = 1e9 + 7
