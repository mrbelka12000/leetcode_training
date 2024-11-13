package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	var (
		result int64
		n      = len(nums)
	)
	// fmt.Println(nums)

	for i := 0; i < n; i++ {

		l, r := i+1, n-1
		for l <= r {
			mid := (l + r) / 2
			sum := nums[i] + nums[mid]

			if sum > upper {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		ll, rr := i+1, n-1
		for ll <= rr {
			mid := (ll + rr) / 2
			sum := nums[i] + nums[mid]

			if sum >= lower {
				rr = mid - 1
			} else {
				ll = mid + 1
			}
		}

		result += int64(r - rr)
	}

	return result
}
