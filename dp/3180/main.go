package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTotalReward([]int{1, 6, 4, 3, 2}))
}

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	dp := getArr(len(rewardValues))

	return runner(rewardValues, dp, 0, 0)
}

func runner(nums []int, dp [][]int, sum, i int) int {
	if i >= len(nums) {
		return sum
	}
	if dp[i][sum] != -1 {
		return dp[i][sum]
	}

	ex := runner(nums, dp, sum, i+1)
	var in int
	if nums[i] > sum {
		in = runner(nums, dp, sum+nums[i], i+1)
	}
	dp[i][sum] = max(in, ex)
	return dp[i][sum]
}

func getArr(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 4001)
		for j := range arr[i] {
			arr[i][j] = -1
		}
	}
	return arr
}
