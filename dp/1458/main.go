package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return solve(nums1, nums2, 0, 0, dp)
}

func solve(n1, n2 []int, i, j int, dp [][]int) int {

	if i >= len(n1) || j >= len(n2) {
		return -math.MaxInt32
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	sumBoth := solve(n1, n2, i+1, j+1, dp)
	sumL := solve(n1, n2, i+1, j, dp)
	sumR := solve(n1, n2, i, j+1, dp)

	result := max(
		sumBoth+(n1[i]*n2[j]),
		sumL,
		sumR,
		n1[i]*n2[j],
	)
	dp[i][j] = result

	return result
}
