package main

import "fmt"

func main() {
	//fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
}

func findBall(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		if val, ok := runner(grid, 0, i); ok {
			ans[i] = val
			continue
		}
		ans[i] = -1
	}

	return ans
}

func runner(grid [][]int, x, y int) (int, bool) {
	if y < 0 || y >= len(grid[0]) {
		return 0, false
	}

	if x >= len(grid) {
		return y, true
	}

	switch grid[x][y] {
	case 1:
		if y != len(grid[x])-1 {
			if grid[x][y+1] == -1 {
				return 0, false
			}
		}

		if val, ok := runner(grid, x+1, y+1); ok {
			return val, true
		}
	case -1:

		if y != 0 {
			if grid[x][y-1] == 1 {
				return 0, false
			}
		}

		if val, ok := runner(grid, x+1, y-1); ok {
			return val, true
		}
	}

	return 0, false
}
