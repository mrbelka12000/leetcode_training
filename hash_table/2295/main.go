package main

import "fmt"

func main() {
	fmt.Println(arrayChange([]int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}))
}

func arrayChange(nums []int, operations [][]int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for _, v := range operations {
		ind := mp[v[0]]
		nums[ind] = v[1]
		mp[v[1]] = ind
	}

	return nums
}
