package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfWays(3))
}

func numOfWays(n int) int {
	dp = make(map[[4]int]int)
	return solve(n, 0, 0, 0, 0)
}

var (
	dp = make(map[[4]int]int)
)

func solve(n, i, col1, col2, col3 int) int {
	if i == n {
		return 1
	}

	if val, ok := dp[[4]int{i, col1, col2, col3}]; ok {
		return val
	}
	var count int
	for c1 := 1; c1 <= 3; c1++ {
		if c1 == col1 {
			continue
		}

		for c2 := 1; c2 <= 3; c2++ {
			if c2 == col2 || c2 == c1 {
				continue
			}

			for c3 := 1; c3 <= 3; c3++ {
				if c3 == col3 || c3 == c2 {
					continue
				}

				count = count + (solve(n, i+1, c1, c2, c3) % MOD)
			}
		}
	}

	dp[[4]int{i, col1, col2, col3}] = count

	return count
}

const MOD = 1000000007
