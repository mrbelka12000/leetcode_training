package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 1}))
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

	var m2 = map[int]bool{}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if m[i] || m2[n] {
			continue
		}
		m[i] = true
		m2[n] = true
		recursion(append(cur, n), nums, m, results)
		m[i] = false
	}
	return
}

// 6
