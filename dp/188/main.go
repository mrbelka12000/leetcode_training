package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func maxProfit(k int, prices []int) int {

	dp := make([][][2]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	var solver func(i, remaining int, holding int) int

	solver = func(i, remaining int, holding int) int {
		if i == len(prices) || remaining == 0 {
			return 0
		}

		if dp[i][remaining][holding] != 0 {
			return dp[i][remaining][holding]
		}

		var (
			doNothing = solver(i+1, remaining, holding)
			doSmth    int
		)
		if holding == 0 {
			doSmth = prices[i] + (solver(i+1, remaining-1, 1))
		} else {
			doSmth = -prices[i] + (solver(i+1, remaining, 0))
		}

		dp[i][remaining][holding] = max(doSmth, doNothing)

		return max(doSmth, doNothing)
	}

	return solver(0, k, 1)
}
