package main

import "fmt"

func main() {
	fmt.Println(minimumSwaps([]int{3, 4, 5, 5, 3, 1}))
}

func minimumSwaps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var (
		maxVal, maxInd int
		minVal         = nums[0]
		minInd         int
	)

	for i, v := range nums {
		if v >= maxVal {
			maxVal = v
			maxInd = i
		}
		if v < minVal {
			minVal = v
			minInd = i
		}
	}

	if maxInd < minInd {
		minInd--
	}

	return (len(nums) - 1 - maxInd) + minInd
}
