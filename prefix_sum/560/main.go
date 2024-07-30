package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func subarraySum(nums []int, k int) int {
	sums := make(map[int]int)
	sums[0] = 1

	var result, sum int

	for _, v := range nums {
		sum += v
		result += sums[sum-k]
		sums[sum]++
	}

	return result
}
