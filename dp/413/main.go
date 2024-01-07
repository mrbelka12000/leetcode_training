package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 8, 9, 10}))
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	var result int
	for i := 2; i < len(nums); i++ {
		var check bool
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			//fmt.Println(nums[i-2], nums[i-1], nums[i], "first", i)
			result++
			check = true
		}

		if !check {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[j-1] != nums[j-1]-nums[j-2] {
				break
			}
			//fmt.Println(nums[j-2], nums[j-1], nums[j], "break", i)
			result++
		}
		//fmt.Println(dp, i)
	}

	//fmt.Println(dp)
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 1,2,3,4,5,6
// RES:
// [1,2,3] , [1,2,3,4], [1,2,3,4,5], [1,2,3,4,5,6], [2,3,4], [2,3,4,5], [2,3,4,5,6], [3,4,5], [3,4,5,6], [4,5,6]
