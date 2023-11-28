package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dividePlayers([]int{3, 2, 5, 1, 3, 4}))
}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	l, r := 0, len(skill)-1

	var arr []int64
	var prev int

	for l < r {
		num := skill[l] * skill[r]
		if len(arr) != 0 {
			// fmt.Println(arr, num,prev)
			if prev != skill[l]+skill[r] {
				return -1
			}
		}
		arr = append(arr, int64(num))
		prev = skill[l] + skill[r]
		l++
		r--
	}

	var result int64

	for _, v := range arr {
		result += v
	}

	return result
}
