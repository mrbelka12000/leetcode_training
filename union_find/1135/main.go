package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumCost(3, [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}))
}

func minimumCost(n int, connections [][]int) int {

	uf := make([]int, n+1)
	for i := 1; i < len(uf); i++ {
		uf[i] = i
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	var result int
	for _, v := range connections {
		x := find(uf, v[0])
		y := find(uf, v[1])

		if x == y {
			continue
		}

		result += v[2]
		union(uf, x, y)
	}

	mp := make(map[int]bool)

	for _, v := range uf[1:] {
		mp[find(uf, v)] = true
	}

	if len(mp) != 1 {
		return -1
	}

	return result
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
