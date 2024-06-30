package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDif := math.MaxInt32

	for i := 0; i < len(arr)-1; i++ {
		minDif = min(abs(arr[i]-arr[i+1]), minDif)
	}

	var result [][]int
	// fmt.Println(minDif)
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) == minDif {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
