package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))
}

func matchPlayersAndTrainers(p []int, t []int) int {
	sort.Ints(p)
	sort.Ints(t)

	var l, r, result int

	for l < len(p) && r < len(t) {
		if p[l] > t[len(t)-1] {
			break
		}

		if p[l] <= t[r] {
			result++
			l++
			r++
		} else if p[l] > t[r] {
			r++
		}

	}

	return result
}
