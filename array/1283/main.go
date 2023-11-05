package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
}

func smallestDivisor(nums []int, th int) int {
	var max int

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return bSearch(nums, th, max)
}

func bSearch(nums []int, th, max int) int {
	var result int
	l, r := 1, max
	for l <= r {
		del := (l + r) / 2

		sum := getSum(nums, float64(del))
		if sum > th {
			l = del + 1
		} else {
			r = del - 1
			result = del
		}
	}

	return result
}

func getSum(nums []int, del float64) (sum int) {
	for _, v := range nums {
		sum += int(math.Ceil(float64(v) / del))
	}
	return sum
}
