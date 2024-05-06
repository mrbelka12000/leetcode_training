package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDistance("horse", "ros"))

}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := range word2 {
		dp[len(word1)][i] = len(word2) - i
	}

	for i := range word1 {
		dp[i][len(word2)] = len(word1) - i
	}

	dp[len(word1)][len(word2)] = 0

	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1])
			}
		}
	}

	return dp[0][0]
}
