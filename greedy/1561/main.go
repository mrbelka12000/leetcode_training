package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
}

func maxCoins(piles []int) int {
	sort.Ints(piles)

	var result int

	for len(piles) >= 3 {
		result += piles[len(piles)-2]
		piles = piles[1:]
		piles = piles[:len(piles)-2]
	}

	return result
}
