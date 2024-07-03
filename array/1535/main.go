package main

import (
	"fmt"
)

func main() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
}

func getWinner(arr []int, k int) int {
	if k >= len(arr) {
		return getMax(arr)
	}
	cpK := k
	var last int

	for k > 0 {
		// fmt.Println(arr, "before")
		first := arr[0]
		arr = arr[1:]
		if first > arr[0] {
			k--
			last = first
			arr = append(arr, arr[0])
			arr[0] = first
		} else {
			k = cpK - 1
			last = arr[0]
			arr = append(arr, first)
		}
		// fmt.Println(arr, "after")
	}

	return last
}

func getMax(arr []int) (result int) {
	for _, v := range arr {
		result = max(result, v)
	}
	return result
}
