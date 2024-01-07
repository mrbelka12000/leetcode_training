package main

import "fmt"

func main() {

	//fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 2, 8, 1, 10}))
	//fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
	fmt.Println(numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}))
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var result int
	dp := make([]map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
	}

	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			dp[i][nums[i]-nums[j]] += dp[j][nums[i]-nums[j]] + 1
		}
	}

	for i := 0; i < len(dp); i++ {
		//fmt.Println(dp[i])
		for _, v := range dp[i] {
			result += v
		}
	}

	return result - (len(nums) * (len(nums) - 1) / 2)
}
