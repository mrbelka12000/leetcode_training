package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

func deleteAndEarn(nums []int) int {
	freq := make(map[int]int)
	sort.Ints(nums)

	for _, v := range nums {
		freq[v]++
	}

	var arr []int

	used := make(map[int]bool)

	for _, v := range nums {
		if !used[v] {
			arr = append(arr, v)
			used[v] = true
		}
	}

	if len(arr) == 1 {
		return freq[arr[0]] * arr[0]
	}
	dp := make(map[int]int)
	dp[0] = freq[arr[0]] * arr[0]
	if arr[0]+1 != arr[1] {
		dp[1] = freq[arr[1]]*arr[1] + dp[0]
	}

	dp[1] = max(dp[0], arr[1]*freq[arr[1]], dp[1])

	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-2]+freq[arr[i]]*arr[i], dp[i-1])
		if arr[i-1]+1 != arr[i] {
			dp[i] = max(dp[i], dp[i-1]+freq[arr[i]]*arr[i])
		}
	}

	return dp[len(arr)-1]
}
