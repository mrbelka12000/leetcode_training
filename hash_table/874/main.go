package main

import (
	"fmt"
)

func main() {
	fmt.Println(robotSim([]int{-2, -1, -2, 3, 7}, [][]int{{1, -3}, {2, -3}, {4, 0}, {-2, 5}, {-5, 2}, {0, 0}, {4, -4}, {-2, -5}, {-1, -2}, {0, 2}}))
}

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	f := 0
	mp := make(map[[2]int]bool)
	for _, v := range obstacles {
		mp[[2]int{v[0], v[1]}] = true
	}

	var (
		result, x, y int
	)

	for _, v := range commands {
		if v == -2 {
			f--
		}
		if v == -1 {
			f++
		}
		if f == 4 {
			f = 0
		}
		if f < 0 {
			f = 3
		}

		for i := 0; i < v; i++ {
			if mp[[2]int{x + dirs[f][0], y + dirs[f][1]}] {
				break
			}
			x += dirs[f][0]
			y += dirs[f][1]
		}
		result = max(result, dist(x, y))
	}

	return result
}

func dist(x, y int) int {
	return x*x + y*y
}
