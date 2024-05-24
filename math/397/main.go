package main

import "fmt"

func main() {
	fmt.Println(integerReplacement(15))
}

func integerReplacement(n int) int {
	result := runner(n, 0)

	// fmt.Println(memo)

	return result - 1
}

var memo = make(map[int]int)

func runner(n, dep int) int {
	if n == 1 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	if n%2 == 0 {
		memo[n] = minVal(memo[n], runner(n/2, dep+1)) + 1
	} else {
		memo[n] = minVal(memo[n], runner(n+1, dep+1)) + 1
		memo[n] = minVal(memo[n], runner(n-1, dep+1)) + 1
	}

	return memo[n]
}

func minVal(a, b int) int {
	if a == 0 {
		return b
	}
	return min(a, b)
}
