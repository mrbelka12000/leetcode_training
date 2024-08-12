package main

import (
	"fmt"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var (
		result  []string
		last    = nums[0]
		lastInd int
	)

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			if i-lastInd == 1 {
				result = append(result, fmt.Sprint(last))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", last, nums[i-1]))
			}
			lastInd = i
			last = nums[i]
		}
	}

	if lastInd == len(nums)-1 {
		result = append(result, fmt.Sprint(last))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", last, nums[len(nums)-1]))
	}

	return result
}
