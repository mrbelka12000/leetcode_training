package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
}

func minCost(costs [][]int) int {
	dp := make([]int, 3)
	for i, v := range costs[0] {
		dp[i] = v
	}

	for i := 1; i < len(costs); i++ {
		newDp := make([]int, 3)
		for j, v := range costs[i] {
			for jj, val := range dp {
				if j == jj {
					continue
				}
				if newDp[j] == 0 || newDp[j] > val+v {
					newDp[j] = val + v
				}
			}
		}
		dp = newDp
	}

	minVal := dp[0]
	for _, v := range dp {
		minVal = min(minVal, v)
	}

	return minVal
}
