package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(successfulPairs([]int{5}, []int{1, 2, 3, 4, 5}, 7))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, len(spells))

	sort.Ints(potions)

	for i := 0; i < len(spells); i++ {
		ind := bSearch(potions, spells[i], int(success))
		// fmt.Println(ind)
		result[i] = len(potions) - ind
	}

	return result
}

func bSearch(nums []int, multy, success int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		// fmt.Println(mid, nums[mid])
		if nums[mid]*multy >= success {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// fmt.Println(multy, r)
	return r
}
