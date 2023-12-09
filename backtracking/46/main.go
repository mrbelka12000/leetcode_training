package main

import "fmt"

func main() {

	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var m = make(map[int]bool, len(nums))
	var results [][]int
	recursion(nil, nums, m, &results)
	return results
}

func recursion(cur, nums []int, m map[int]bool, results *[][]int) {
	if len(nums) == len(cur) {
		var t = make([]int, len(cur))
		copy(t, cur)
		*results = append(*results, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if m[n] {
			continue
		}
		m[n] = true
		recursion(append(cur, n), nums, m, results)
		m[n] = false
	}
	return
}

// 6
