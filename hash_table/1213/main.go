package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(arraysIntersection(
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 5, 7, 9},
		[]int{1, 3, 4, 5, 8},
	))
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	mp := make(map[int]int)

	var (
		minVal = math.MaxInt32
		maxVal int
	)
	for _, v := range arr1 {
		mp[v]++
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}

	for _, v := range arr2 {
		mp[v]++
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}

	for _, v := range arr3 {
		mp[v]++
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}

	var result []int

	for i := minVal; i <= maxVal; i++ {
		if mp[i] == 3 {
			result = append(result, i)
		}
	}

	return result
}
