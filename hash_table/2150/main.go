package main

import "fmt"

func main() {
	fmt.Println(findLonely([]int{10, 6, 5, 8}))
}

func findLonely(nums []int) []int {
	mp := make(map[int]int)

	for _, v := range nums {
		mp[v]++
	}

	var result []int

	// fmt.Println(mp)
	for i := 0; i < len(nums); i++ {
		val := mp[nums[i]]
		if val != 1 {
			continue
		}

		if _, ok := mp[nums[i]+1]; ok {
			continue
		}

		if _, ok := mp[nums[i]-1]; ok {
			continue
		}
		result = append(result, nums[i])
	}

	return result
}
