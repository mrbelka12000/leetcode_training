package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumCardPickup([]int{3, 4, 2, 3, 4, 7}))
}

func minimumCardPickup(cards []int) int {
	result := math.MaxInt32
	mp := make(map[int]int)

	for i, v := range cards {
		if val, ok := mp[v]; ok {
			result = min(result, i-val+1)
		}
		mp[v] = i
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}
