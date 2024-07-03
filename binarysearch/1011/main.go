package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func shipWithinDays(weights []int, days int) int {
	l, r := 0, 100000000
	for r-l > 1 {
		mid := (l + r) / 2
		if getCost(weights, mid) > days {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}

func getCost(w []int, c int) int {
	tmp := c
	result := 1
	for _, v := range w {
		if v > tmp {
			return math.MaxInt32
		}

		c -= v
		if c < 0 {
			c = tmp - v
			result++
		}
	}

	return result
}
