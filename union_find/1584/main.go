package main

import "sort"

func main() {

}

func minCostConnectPoints(p [][]int) int {
	uf := make([]int, len(p))
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}

	var graph [][3]int

	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			dist := getDist(p[i], p[j])
			graph = append(graph, [3]int{i, j, dist})
		}
	}

	sort.Slice(graph, func(i, j int) bool {
		return graph[i][2] < graph[j][2]
	})

	var result int
	// mp := make(map[int]int)
	var count int
	for _, v := range graph {
		if count == len(p)-1 {
			break
		}

		xPos := find(uf, v[0])
		yPos := find(uf, v[1])
		if xPos == yPos {
			continue
		}

		result += v[2]
		count++
		union(uf, xPos, yPos)
		// fmt.Println(uf, v, xPos, yPos)
	}

	// fmt.Println(mp)

	// fmt.Println(uf)

	return result
}

func union(uf []int, x, y int) {
	for i, v := range uf {
		if v == y {
			uf[i] = x
			// fmt.Println(x,y, i)
		}
	}

	uf[y] = x
}

func find(uf []int, x int) int {
	if uf[x] != x {
		uf[x] = find(uf, uf[x])
	}
	return uf[x]
}

func getDist(x, y []int) int {
	val := abs(x[0]-y[0]) + abs(x[1]-y[1])
	return val
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
