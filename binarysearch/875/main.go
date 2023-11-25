package main

import (
	"fmt"
)

func main() {
	//fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	//fmt.Println(minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 823855818))
	//fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
	fmt.Println(minEatingSpeed([]int{1, 1, 1, 999999999}, 10))
}

func minEatingSpeed(piles []int, h int) int {

	var (
		hours  int
		max    = getMax(piles)
		result = max
		min    = 1
		del    = (max + min) / 2
	)

	// 823855818
	// 771945402
	for min <= max {
		hours = calculateHours(piles, del)

		if hours > h {
			min = del + 1
		} else if hours <= h {
			max = del - 1
			result = minVal(result, del)
		}

		del = (max + min) / 2
	}

	return result
}

func minVal(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func calculateHours(piles []int, del int) (hours int) {
	for _, v := range piles {
		hours += v / del
		if v%del > 0 {
			hours++
		}
		//hours += int(math.Ceil(float64(v) / float64(del)))
	}

	return hours
}

func getMax(piles []int) (max int) {

	for _, v := range piles {
		if v > max {
			max = v
		}
	}

	return max
}
