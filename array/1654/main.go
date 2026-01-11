package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	fmt.Println(minimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11))
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	visited := make(map[int]struct{}, len(forbidden))
	var maxForb int
	for i := 0; i < len(forbidden); i++ {
		visited[forbidden[i]] = struct{}{}
		maxForb = max(maxForb, forbidden[i])
	}

	bound := max(maxForb+a+b, x+b)

	var result int

	type check struct {
		Pos  int
		Back bool
	}

	q := []check{{Pos: 0}}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			pos := q[0]
			q = q[1:]

			curr := pos.Pos

			if curr == x {
				return result
			}

			if curr > bound {
				continue
			}
			if curr-b > 0 && !pos.Back {
				if _, ok := visited[curr-b]; !ok {
					visited[curr-b] = struct{}{}
					q = append(q, check{Pos: curr - b, Back: true})
				}
			}
			if _, ok := visited[curr+a]; !ok {
				visited[curr+a] = struct{}{}
				q = append(q, check{Pos: curr + a, Back: false})
			}
		}
		result++
	}

	return -1
}
