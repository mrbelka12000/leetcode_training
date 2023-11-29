package main

import "fmt"

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
}

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i-k < 0 {
			result[i] = -1
			continue
		}
		if i+k >= len(nums) {
			result[i] = -1
			continue
		}
		// fmt.Println(nums[i-k:i+k])
		result[i] = getAvg(nums[i-k : i+k+1])
	}

	return result
}

func getAvg(arr []int) int {
	var sum int

	for _, v := range arr {
		sum += v
	}

	return sum / len(arr)
}
