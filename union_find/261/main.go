package main

import (
	"fmt"
)

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {2, 4}}))
}

func validTree(n int, edges [][]int) bool {
	uf := make([]int, n)
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}

	for _, v := range edges {
		x := find(uf, v[0])
		y := find(uf, v[1])
		if x == y {
			return false
		}

		union(uf, x, y)
	}

	check := make(map[int]bool)
	for _, v := range uf {
		check[find(uf, v)] = true
	}
	return len(check) == 1
}

func union(uf []int, x, y int) {
	uf[y] = x
}

func find(uf []int, x int) int {
	for uf[x] != x {
		x = uf[x]
	}
	return x
}
