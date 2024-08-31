package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumCosts([]int{1, 6, 9, 5}, []int{5, 2, 3, 10}, 8))
}

func minimumCosts(reg []int, exp []int, eCost int) []int64 {
	n := len(reg)
	regular := make(map[int]int)
	express := make(map[int]int)

	regular[-1] = 0
	express[-1] = eCost

	result := make([]int64, n)
	for i := 0; i < n; i++ {
		regular[i] = min(express[i-1], regular[i-1]) + reg[i]
		express[i] = min(express[i-1], regular[i-1]+eCost) + exp[i]
		result[i] = int64(min(regular[i], express[i]))
	}

	return result
}
