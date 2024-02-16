package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
}

func frequencySort(nums []int) []int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	var result []int

	var check [][2]int

	for k, v := range freq {
		check = append(check, [2]int{k, v})
	}

	sort.Slice(check, func(i, j int) bool {
		if check[i][1] == check[j][1] {
			return check[i][0] > check[j][0]
		}
		return check[i][1] < check[j][1]
	})

	// fmt.Println(check,freq)

	for i := 0; i < len(check); i++ {
		for j := 0; j < check[i][1]; j++ {
			result = append(result, check[i][0])
		}
	}

	return result
}
