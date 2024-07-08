package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func maxProfitAssignment(dif []int, profit []int, worker []int) int {
	store := make([][]int, len(dif))
	for i := 0; i < len(dif); i++ {
		store[i] = []int{dif[i], profit[i]}
	}

	sort.Slice(store, func(i, j int) bool {
		if store[i][0] == store[j][0] {
			return store[i][1] < store[j][1]
		}
		return store[i][0] < store[j][0]
	})

	var maxVal int
	for i, v := range store {
		if maxVal > v[1] {
			store[i][1] = maxVal
		}
		maxVal = max(maxVal, v[1])
	}

	var result int
	for _, v := range worker {
		l, r := -1, len(store)
		for r-l > 1 {
			mid := (r + l) / 2

			if store[mid][0] > v {
				r = mid
			} else {
				l = mid
			}
		}

		if l != -1 {
			result += store[l][1]
		}
	}

	return result
}
