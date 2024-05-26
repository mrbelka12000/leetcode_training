package main

import "fmt"

func main() {

}

func numEnclaves(grid [][]int) int {
	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				result += runner(grid, i, j)
			}
		}
	}
	return result
}

func runner(grid [][]int, x, y int) int {
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	q := [][2]int{{x, y}}
	var (
		check bool
		area  int
	)

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if grid[n[0]][n[1]] == 0 {
			continue
		}

		grid[n[0]][n[1]] = 0
		area++
		for _, dir := range dirs {
			xx, yy := n[0]+dir[0], n[1]+dir[1]

			if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[xx]) {
				check = true
				continue
			}

			if grid[xx][yy] == 0 {
				continue
			}

			q = append(q, [2]int{xx, yy})
		}
	}

	if check {
		return 0
	}

	return area
}

func Print(grid [][]int) {
	for _, v := range grid {
		fmt.Println(v)
	}
	fmt.Println()
}

/*
[[0,0,0,1,1,1,0,1,1,1,1,1,0,0,0]
,[1,1,1,1,0,0,0,1,1,0,0,0,1,1,1]
,[1,1,1,0,0,1,0,1,1,1,0,0,0,1,1]
,[1,1,0,1,0,1,1,0,0,0,1,1,0,1,0]
,[1,1,1,1,0,0,0,1,1,1,0,0,0,1,1]
,[1,0,1,1,0,0,1,1,1,1,1,1,0,0,0]
,[0,1,0,0,1,1,1,1,0,0,1,1,1,0,0]
,[0,0,1,0,0,0,0,1,1,0,0,1,0,0,0]
,[1,0,1,0,0,1,0,0,0,0,0,0,1,0,1]
,[1,1,1,0,1,0,1,0,1,1,1,0,0,1,0]]

[0 0 0 0 0 0 0 1 1 1 1 1 0 0 0]
[0 0 0 0 0 0 0 1 1 0 0 0 1 1 1]
[0 0 0 0 0 1 0 1 1 1 0 0 0 1 1]
[0 0 0 0 0 1 1 0 0 0 1 1 0 1 0]
[0 0 0 0 0 0 0 1 1 1 0 0 0 1 1]
[0 0 0 0 0 0 1 1 1 1 1 1 0 0 0]
[0 1 0 0 1 1 1 1 0 0 1 1 1 0 0]
[0 0 1 0 0 0 0 1 1 0 0 1 0 0 0]
[1 0 1 0 0 1 0 0 0 0 0 0 1 0 1]
[1 1 1 0 1 0 1 0 1 1 1 0 0 1 0]
*/
