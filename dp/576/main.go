package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
	//fmt.Println(findPaths(8, 7, 15, 4, 1))
	//fmt.Println(findPaths(1, 3, 3, 0, 1))
}

var (
	mod  = int(math.Pow(10, 9) + 7)
	dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var runner func(i, j, moves int) int
	cache := make(map[[3]int]int)
	runner = func(i, j, moves int) int {
		key := [3]int{i, j, moves}
		if val, ok := cache[key]; ok {
			return val
		}
		if i >= 0 && i < m && j >= 0 && j < n {
			if moves == 0 {
				cache[key] = 0
				return cache[key]
			}
			for _, dir := range dirs {
				cache[key] += runner(i+dir[0], j+dir[1], moves-1) % mod
			}
			return cache[key]
		} else {
			cache[key] = 1
			return cache[key]
		}
	}

	return runner(startRow, startColumn, maxMove) % mod
}
