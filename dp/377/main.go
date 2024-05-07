package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

func combinationSum4(nums []int, target int) int {
	return runner(nums, target, make(map[int]int, 10))
}

func runner(nums []int, target int, mp map[int]int) int {
	if target == 0 {
		return 1
	}

	if val, ok := mp[target]; ok {
		return val
	}

	mp[target] = 0

	for _, v := range nums {
		if target-v < 0 {
			continue
		}
		mp[target] += runner(nums, target-v, mp)
	}

	return mp[target]
}
