package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{1, -1, 2, 4}, 2))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	// fmt.Println(nums)
	var result int = 12342134123

	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1

		for l < r {
			num := nums[i] + nums[l] + nums[r]
			if num < target {
				l++
			} else if num > target {
				r--
			} else {
				result = num
				break
			}
			// fmt.Println(num,i, l, r)

			if abs(num-target) < abs(result-target) {
				result = num
			}
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
