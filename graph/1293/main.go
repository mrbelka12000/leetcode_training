package main

import "fmt"

func main() {
	fmt.Println(shortestPath(
		[][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0}},
		2,
	))
}

func shortestPath(grid [][]int, k int) int {

	var (
		q = [][3]int{
			{0, 0, k},
		}
		visited = make(map[[3]int]bool)
		dirs    = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		cost    int
	)

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			x, y, r := n[0], n[1], n[2]
			if x == len(grid)-1 && y == len(grid[0])-1 {
				return cost
			}
			for _, dir := range dirs {
				xx, yy := x+dir[0], y+dir[1]
				if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[0]) {
					continue
				}

				rr := r
				if grid[xx][yy] == 1 {
					if rr == 0 {
						continue
					}
					rr--
				}
				if !visited[[3]int{xx, yy, rr}] {
					visited[[3]int{xx, yy, rr}] = true
					q = append(q, [3]int{xx, yy, rr})
				}

			}
		}

		cost++
	}
	return -1
}

/*
[[0,1,0,0,0,1,0,0]
,[0,1,0,1,0,1,0,1]
,[0,0,0,1,0,0,1,0]]
*/
