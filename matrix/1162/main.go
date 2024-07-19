package main

func main() {

}

func maxDistance(grid [][]int) int {
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var q [][2]int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	var (
		dist   int
		exists bool
	)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[i]
			for _, dir := range dirs {
				xx, yy := n[0]+dir[0], n[1]+dir[1]
				if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[xx]) || grid[xx][yy] == 1 {
					continue
				}
				exists = true
				grid[xx][yy] = 1
				q = append(q, [2]int{xx, yy})
			}
		}

		dist++
		q = q[size:]
	}

	if !exists {
		return -1
	}

	return dist - 1
}
