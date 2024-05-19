package main

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))

}

func isBipartite(graph [][]int) bool {

	arr := make([]int, len(graph))

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i, v := range graph {
		if len(v) == 0 {
			continue
		}

		x := find(arr, i)
		for _, edge := range v {
			y := find(arr, edge)
			if x == y {
				return false
			}
			union(arr, find(arr, v[0]), y)
		}
	}

	return true
}

func find(prov []int, x int) int {
	if prov[x] != x {
		prov[x] = find(prov, prov[x])
	}

	return prov[x]
}

func union(prov []int, x, y int) {
	prov[x] = y
}
