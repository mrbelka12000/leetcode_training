package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}

func subarraysDivByK(nums []int, k int) int {
	store := make(map[int]int)
	var result, sum int
	store[0] = 1

	for _, v := range nums {
		sum += v

		rem := sum % k

		if rem < 0 {
			rem += k
		}

		result += store[rem]

		store[rem]++
	}

	return result
}
