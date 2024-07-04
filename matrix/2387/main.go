package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func matrixMedian(grid [][]int) int {
	var arr []int

	for _, v := range grid {
		for _, val := range v {
			arr = append(arr, val)
		}
	}

	sort.Ints(arr)

	return arr[len(arr)/2]
}
