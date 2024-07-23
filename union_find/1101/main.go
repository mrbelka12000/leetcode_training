package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestAcq([][]int{{9, 0, 3}, {0, 2, 7}, {12, 3, 1}, {5, 5, 2}, {3, 4, 5}, {1, 5, 0}, {6, 2, 4}, {2, 5, 3}, {7, 7, 3}}, 8))
}

func earliestAcq(logs [][]int, n int) int {
	uf := make([]int, n)
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	for _, v := range logs {
		x := find(uf, v[1])
		y := find(uf, v[2])
		if x == y {
			continue
		}

		union(uf, x, y)
		var arr [101]int
		for _, v := range uf {
			arr[find(uf, v)]++
		}
		if arr[x] == n {
			return v[0]
		}
	}

	return -1
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
