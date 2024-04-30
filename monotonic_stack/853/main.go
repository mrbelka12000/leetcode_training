package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{6, 8}, []int{3, 2}))
}

func carFleet(target int, pos []int, speed []int) int {
	mp := make(map[int]int)
	for i, v := range pos {
		mp[v] = i
	}

	sort.Ints(pos)

	var stack []float64

	for i := 0; i < len(pos); i++ {
		curr := float64(speed[mp[pos[i]]]) / float64(target-pos[i])
		for len(stack) > 0 && stack[len(stack)-1] >= curr {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, curr)
	}

	return len(stack)
}
