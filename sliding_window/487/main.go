package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	var result, start, o, z int

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v == 0 {
			z++
		} else {
			o--
		}

		for z > 1 {
			if nums[start] == 0 {
				z--
			} else {
				o--
			}
			start++
		}

		result = max(result, i-start+1)
	}

	return result
}
