package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
}

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 1 {
		return 1
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	n := len(pairs)

	dp := make([][2]int, n)
	dp[0] = [2]int{1, pairs[0][1]}
	for i := 1; i < n; i++ {
		last := dp[i-1]
		if last[1] < pairs[i][0] {
			dp[i] = [2]int{last[0] + 1, pairs[i][1]}
		} else {
			dp[i] = [2]int{last[0], min(pairs[i][1], last[1])}
		}
	}

	return dp[n-1][0]
}
