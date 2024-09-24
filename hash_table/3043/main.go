package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	mp := make(map[int]int)

	for _, v := range arr2 {
		size := len(fmt.Sprint(v))

		for v > 0 {
			mp[v] = size
			v /= 10
			size--
		}
	}

	var result int
	for _, v := range arr1 {
		for v > 0 {
			result = max(result, mp[v])
			v /= 10
		}
	}

	// fmt.Println(mp)
	return result
}
