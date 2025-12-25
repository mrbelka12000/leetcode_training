package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
}

func maximumHappinessSum(h []int, k int) int64 {
	sort.Slice(h, func(i, j int) bool {
		return h[i] > h[j]
	})

	var (
		result int64
		decr   int
	)

	for i := 0; i < len(h) && k > 0; i++ {
		val := h[i] - decr
		if val > 0 {
			result += int64(val)
		}

		decr++
		k--
	}

	return result
}
