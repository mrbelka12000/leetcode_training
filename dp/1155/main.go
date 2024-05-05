package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numRollsToTarget(30, 30, 500))
}

func numRollsToTarget(n int, k int, target int) int {
	return solver(n, k, target, make(map[[2]int]int)) % mod
}

var mod = int(math.Pow(10, 9) + 7)

func solver(n int, faces int, target int, memo map[[2]int]int) int {
	if target == 0 && n == 0 {
		return 1
	}

	if val, ok := memo[[2]int{target, n}]; ok {
		return val
	}

	for i := 1; i <= faces; i++ {
		v := i
		if n != 0 && target-v >= 0 {
			memo[[2]int{target, n}] = (memo[[2]int{target, n}] + solver(n-1, faces, target-v, memo)%mod)
		}
	}

	return memo[[2]int{target, n}]
}
