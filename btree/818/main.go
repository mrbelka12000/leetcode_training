package main

import (
	"fmt"
)

func main() {
	fmt.Println(racecar(21))
}

func racecar(target int) int {
	var result int

	q := [][2]int{{0, 1}}
	visited := make(map[[2]int]bool)

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[i]

			if n[0] == target {
				return result
			}

			if n[1] > 0 && !visited[[2]int{n[0], -1}] {
				q = append(q, [2]int{n[0], -1})
				visited[[2]int{n[0], -1}] = true
			}

			if n[1] < 0 && !visited[[2]int{n[0], 1}] {
				q = append(q, [2]int{n[0], 1})
				visited[[2]int{n[0], 1}] = true
			}

			if !visited[[2]int{n[0] + n[1], n[1] * 2}] {
				q = append(q, [2]int{n[0] + n[1], n[1] * 2})
				visited[[2]int{n[0] + n[1], n[1] * 2}] = true
			}
		}

		q = q[size:]
		result++
	}

	return result
}
