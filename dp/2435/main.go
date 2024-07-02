package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func numberOfPaths(grid [][]int, k int) int {
	return dfs(grid, 0, 0, 0, k, make(map[[3]int]int, len(grid)*len(grid[0]))) % mod
}

func dfs(grid [][]int, x, y, cost, k int, dp map[[3]int]int) int {
	if x == len(grid)-1 && y == len(grid[x])-1 {
		cost += grid[x][y]
		if cost%k == 0 {
			return 1
		}
		return 0
	}

	if x >= len(grid) || y >= len(grid[x]) {
		return 0
	}

	cost += grid[x][y]
	key := [3]int{x, y, cost % k}
	if val, ok := dp[key]; ok {
		return val
	}

	dp[key] = (dp[key] + dfs(grid, x+1, y, cost, k, dp)) % mod
	dp[key] = (dp[key] + dfs(grid, x, y+1, cost, k, dp)) % mod

	return dp[key] % mod
}

const mod = 1e9 + 7
