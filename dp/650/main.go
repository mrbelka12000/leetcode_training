package main

import "fmt"

func main() {
	fmt.Println(minSteps(5))
}

func minSteps(n int) int {
	dp := make(map[int]int)
	dp[2] = 2

	for i := 3; i <= n; i++ {
		pm := getPrime(i)

		if pm == i {
			dp[i] = pm
		} else {
			dp[i] = dp[pm] + dp[i/pm]
		}
	}

	return dp[n]
}

func getPrime(n int) int {
	res := n

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res = i
			n /= i
		}
	}

	if n > 1 {
		res = max(res, n)
	}

	return res
}
