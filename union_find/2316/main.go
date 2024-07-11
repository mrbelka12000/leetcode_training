package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
}

func countPairs(n int, edges [][]int) int64 {
	uf := make([]int, n)
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}

	for _, v := range edges {
		union(uf, v[0], v[1])
	}

	mp := make([]int, n)
	for _, v := range uf {
		mp[find(uf, v)]++
	}

	var result int64
	for i := 0; i < len(uf); i++ {
		gr := find(uf, uf[i])
		result += int64(n - mp[gr])
	}

	return result / 2
}

func union(uf []int, x, y int) {
	x, y = find(uf, x), find(uf, y)

	uf[y] = x
}

func find(uf []int, x int) int {
	for uf[x] != x {
		x = uf[x]
	}

	return uf[x]
}
