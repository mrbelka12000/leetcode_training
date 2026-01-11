package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return solve(s1, s2, 0, 0, dp)
}

func solve(s1, s2 string, i, j int, dp [][]int) int {
	if i >= len(s1) && j >= len(s2) {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if i >= len(s1) {
		return solve(s1, s2, i, j+1, dp) + int(s2[j])
	}
	if j >= len(s2) {
		return solve(s1, s2, i+1, j, dp) + int(s1[i])
	}

	minVal := math.MaxInt32
	if s1[i] == s2[j] {
		minVal = min(minVal, solve(s1, s2, i+1, j+1, dp))
	} else {
		minVal = min(minVal, solve(s1, s2, i+1, j, dp)+int(s1[i]))
		minVal = min(minVal, solve(s1, s2, i, j+1, dp)+int(s2[j]))
	}

	dp[i][j] = minVal
	return minVal
}
