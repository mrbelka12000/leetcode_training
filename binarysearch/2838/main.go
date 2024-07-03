package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
	comb := make([][2]int, 0, len(monsters))

	for i, v := range monsters {
		comb = append(comb, [2]int{v, coins[i]})
	}
	sort.Slice(comb, func(i, j int) bool {
		return comb[i][0] < comb[j][0]
	})

	ps := make([]int, len(comb))
	ps[0] = comb[0][1]

	for i := 1; i < len(ps); i++ {
		ps[i] += ps[i-1] + comb[i][1]
	}

	result := make([]int64, len(heroes))
	for i := 0; i < len(heroes); i++ {
		pos := getPos(comb, heroes[i])
		if pos != -1 {
			result[i] = int64(ps[pos])
		}
	}

	return result
}

func getPos(arr [][2]int, st int) int {
	l, r := -1, len(arr)
	for r-l > 1 {
		mid := (l + r) / 2
		if arr[mid][0] > st {
			r = mid
		} else {
			l = mid
		}
	}

	return l
}
