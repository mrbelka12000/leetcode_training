package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 1, 3}))
}

func majorityElement(nums []int) []int {
	var result []int

	mp := make(map[int]int)

	for _, v := range nums {
		mp[v]++
	}

	del := len(nums) / 3
	for k, v := range mp {
		if v > del {
			result = append(result, k)
		}
	}

	return result
}
