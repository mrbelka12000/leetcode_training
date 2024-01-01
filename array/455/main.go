package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var result int

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] >= g[i] {
				result++
				s[j] = 0
				break
			}
		}
	}

	return result
}
