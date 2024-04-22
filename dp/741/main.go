package main

import "fmt"

func main() {
	//fmt.Println(cherryPickup([][]int{
	//	{0, 1, -1},
	//	{1, -1, -1},
	//	{1, 1, 1}},
	//))

	fmt.Println(cherryPickup([][]int{
		{0, 1, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{-1, 1, 1, 1, -1},
		{0, 1, 1, 1, 0},
		{1, 0, -1, 0, 0},
	}))

}

func cherryPickup(grid [][]int) int {
	check := grid[len(grid)-1][len(grid[0])-1] == 1
	dfs(grid, right_down, 0, 0)

	if grid[len(grid)-1][len(grid[0])-1] != 2 {
		return 0
	}

	for _, v := range grid {
		fmt.Println(v)
	}
	fmt.Println("--------")
	result := run(grid, right_down, 0, 0)
	if check {
		result++
	}
	for _, v := range grid {
		fmt.Println(v)
	}

	return result
}

func run(grid, dirs [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 0
	}

	if grid[x][y] == -1 {
		return 0
	}

	var count int
	if grid[x][y] == 1 {
		count++
	}
	grid[x][y] = -1

	//fmt.Println(x, y)
	count += run(grid, dirs, x+dirs[0][0], y+dirs[0][1])
	count += run(grid, dirs, x+dirs[1][0], y+dirs[1][1])

	return count
}

// dfs return true if blocks
func dfs(grid, dirs [][]int, x, y int) bool {
	if x == len(grid)-1 && y == len(grid[0])-1 {
		grid[x][y] = 2
		return false
	}

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}

	if grid[x][y] == -1 {
		return true
	}

	//fmt.Println(x, y)
	b1 := dfs(grid, dirs, x+dirs[0][0], y+dirs[0][1])
	b2 := dfs(grid, dirs, x+dirs[1][0], y+dirs[1][1])

	if b1 && b2 {
		grid[x][y] = -1
	}

	return b1 == true && b2 == true
}

var (
	right_down = [][]int{{0, 1}, {1, 0}}
	left_up    = [][]int{{0, -1}, {-1, 0}}
)

/*
[[0,1,-1]
,[0,1,-1]
,[1,-1,-1]
,[1,1,1]]
*/

/*
[[0,-1,-1]
,[0,-1,-1]
,[1,-1,-1]
,[1,1,1]]
*/
