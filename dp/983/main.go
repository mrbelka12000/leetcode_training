package main

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}

func mincostTickets(days []int, costs []int) int {
	dp := make(map[int]int)

	travers := make(map[int]bool)
	for _, v := range days {
		travers[v] = true
	}

	last := days[len(days)-1]

	for i := 1; i <= last; i++ {
		if !travers[i] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = min(
			max(dp[i-1]+costs[0], 0),
			max(dp[i-7]+costs[1], 0),
			max(dp[i-30]+costs[2], 0),
		)
	}

	return dp[last]
}
