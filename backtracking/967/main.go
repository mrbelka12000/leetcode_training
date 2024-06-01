package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numsSameConsecDiff(2, 3))
}

func numsSameConsecDiff(n int, k int) []int {
	var res []int

	check := math.Pow(10, float64(n))
	for i := 1; i <= 9; i++ {
		backtrack(n, k, i, int(check), &res)
	}

	return res
}

func backtrack(n, k, curr, check int, res *[]int) {
	if curr*10 >= check {
		*res = append(*res, curr)
		return
	}

	last := curr % 10
	if last+k < 10 {
		backtrack(n, k, curr*10+last+k, check, res)
	}

	if k != 0 && last-k >= 0 {
		backtrack(n, k, curr*10+last-k, check, res)
	}
}
