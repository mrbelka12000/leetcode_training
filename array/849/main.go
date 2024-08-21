package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	var (
		result, left, count int
		hasLeft             bool
	)

	for i, v := range seats {
		if v == 1 {
			count++
			if hasLeft {
				result = max(result, int((i-left)/2))
			} else {
				result = i
			}
			left = i
			hasLeft = true
		}
	}

	if count == 1 {
		return max(len(seats)-1-left, left)
	}

	return max(len(seats)-1-left, result)
}
