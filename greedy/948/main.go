package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	var result, tmp int
	l, r := 0, len(tokens)-1

	for l <= r {
		result = max(result, tmp)
		if tokens[l] <= power {
			tmp++
			power -= tokens[l]
			l++
			continue
		}
		if tmp == 0 {
			break
		}

		tmp--
		power += tokens[r]
		r--
	}

	return max(result, tmp)
}
