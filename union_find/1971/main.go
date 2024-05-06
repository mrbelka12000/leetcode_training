package main

func main() {

}

func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := make([]int, n)
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}

	for _, v := range edges {
		union(uf, v[0], v[1])
	}

	return find(uf, source) == find(uf, destination)
}

func union(uf []int, x, y int) {
	xPos, yPos := find(uf, x), find(uf, y)
	uf[yPos] = xPos
}

func find(uf []int, x int) int {
	if uf[x] != x {
		uf[x] = find(uf, uf[x])
	}

	return uf[x]
}
