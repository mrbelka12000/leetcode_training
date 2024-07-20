package main

import (
	"fmt"
)

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return runner(amount, coins, dp)
}

func runner(amount int, coins []int, dp [][]int) int {
	if amount == 0 {
		return 1
	}

	if dp[len(coins)][amount] != -1 {
		return dp[len(coins)][amount]
	}
	dp[len(coins)][amount] = 0
	for i, v := range coins {
		if amount >= v {
			dp[len(coins)][amount] += runner(amount-v, coins[i:], dp)
		}
	}

	return dp[len(coins)][amount]
}
