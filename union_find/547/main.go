package main

func main() {

}

func findCircleNum(isConnected [][]int) int {
	prov := make([]int, len(isConnected))
	for i := 0; i < len(prov); i++ {
		prov[i] = i
	}

	// fmt.Println(isConnected)
	// fmt.Println(prov)
	for x, v := range isConnected {
		for y, pos := range v {
			if pos == 1 && x != y {
				// fmt.Println(prov, x, y)
				union(prov, x, y)
				// fmt.Println(prov)
			}
		}
		// fmt.Println()
	}

	mp := make(map[int]bool)

	for _, v := range prov {
		mp[v] = true
	}

	return len(mp)
}

func union(prov []int, x, y int) {
	xPos := find(prov, x)
	yPos := find(prov, y)

	// fmt.Printf("x=%v y=%v xPos=%v, yPos=%v\n", x,y, xPos, yPos)
	for i, v := range prov {
		if v == yPos {
			prov[i] = xPos
		}
	}

	prov[yPos] = xPos
}

func find(prov []int, x int) int {
	if prov[x] == x {
		return prov[x]
	}
	return find(prov, prov[x])
}
