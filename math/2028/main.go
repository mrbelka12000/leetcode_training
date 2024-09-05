package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2))
}

func missingRolls(rolls []int, mean int, n int) []int {
	var (
		sum    int
		result = make([]int, n)
	)

	for _, v := range rolls {
		sum += v
	}

	x := (mean*(len(rolls)+n) - sum) / n
	y := (mean*(len(rolls)+n) - sum) % n

	for i := 0; i < n; i++ {
		result[i] = x
	}

	for i := 0; i < y; i++ {
		result[i]++
	}

	for _, v := range result {
		if v <= 0 || v > 6 {
			return nil
		}
	}

	return result
}
