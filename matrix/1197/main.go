package main

import (
	"fmt"
)

func main() {
	fmt.Println(minKnightMoves(2, 112))
}

func minKnightMoves(x int, y int) int {
	dirs := [8][2]int{{2, 1}, {2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
	q := [][2]int{{0, 0}}

	var result int
	visited := make(map[[2]int]bool)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[i]

			if n[0] == x && n[1] == y {
				return result
			}

			for _, dir := range dirs {
				tmp := [2]int{n[0] + dir[0], n[1] + dir[1]}
				if !visited[tmp] {
					q = append(q, tmp)
					visited[tmp] = true
				}
			}
		}

		result++
		q = q[size:]
	}
	return result
}
