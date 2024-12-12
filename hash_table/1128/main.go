package main

import (
	"fmt"
)

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	mp := make(map[[2]int]int)
	var result int

	for _, v := range dominoes {
		val, ok := mp[[2]int{v[0], v[1]}]
		_, ok1 := mp[[2]int{v[1], v[0]}]
		if ok || ok1 {
			result += val
		}
		mp[[2]int{v[0], v[1]}]++
		if v[0] != v[1] {
			mp[[2]int{v[1], v[0]}]++
		}
	}

	return result
}
