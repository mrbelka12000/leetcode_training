package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	//fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	//fmt.Println(jobScheduling([]int{1, 2, 2, 3}, []int{2, 5, 3, 4}, []int{3, 4, 1, 2}))
	fmt.Println(jobScheduling([]int{24, 24, 7, 2, 1, 13, 6, 14, 18, 24}, []int{27, 27, 20, 7, 14, 22, 20, 24, 19, 27}, []int{6, 1, 4, 2, 3, 6, 5, 6, 9, 8}))

	//fmt.Println(jobScheduling(
	//	[]int{6, 15, 7, 11, 1, 3, 16, 2},
	//	[]int{19, 18, 19, 16, 10, 8, 19, 8},
	//	[]int{2, 9, 1, 19, 5, 7, 3, 19}))
}

/*

          4
	-2          5-
  3
-1 2-

	  1
    -2 3-
	      2
		-3 4-
1---2---3---4---5
*/

func jobScheduling(startTime []int, endTime []int, profits []int) int {
	var (
		arr [][3]int
	)
	ln := len(startTime)
	for i := 0; i < ln; i++ {
		arr = append(arr, [3]int{startTime[i], endTime[i], profits[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})

	dp := make([]int, ln+1)
	//dp[0] = arr[0][2]

	// fmt.Println(arr)
	for i := 0; i < ln; i++ {
		fmt.Println("INDEX: ", i)

		nextPos := bSearch(arr, arr[i][0], i)
		//fmt.Println(arr[i], "------TOP------", arr[nextPos])
		// for j := nextPos; j < len(arr); j++ {
		//fmt.Println(arr[j], "INLOOP", arr[j][2], dp[i])
		dp[i+1] = max(dp[i], dp[nextPos]+arr[i][2])
		// result = max(result, dp[j])
		// }

		//fmt.Println(dp, "DPß")

		//for j := i - 1; j >= 0; j-- {
		//	if arr[i][0] >= arr[j][1] {
		//		tmp = max(arr[i][2]+dp[j], tmp)
		////		fmt.Println(arr[j])
		//	}
		//}

		//fmt.Println()
		//dp[nextPos] = max(dp[i]+arr[i][2], dp[nextPos])
	}

	//fmt.Println()
	// fmt.Println(dp)
	//fmt.Println(dp)
	//fmt.Println(arr)
	return dp[len(arr)]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func bSearch(arr [][3]int, target, r int) int {
	l := 0
	for l < r {
		mid := (l + r) / 2

		if arr[mid][1] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
