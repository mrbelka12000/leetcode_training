package main

func main() {

}

func countPaths(grid [][]int) int {
	var result int
	visited = make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			result = (result + dfs(grid, i, j, 12341233124)) % mod
		}
	}

	return result
}

const mod = 10e9 + 7

var (
	dirs    = [4][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	visited [][]int
)

func dfs(grid [][]int, x, y, prev int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || prev <= grid[x][y] {
		return 0
	}

	if visited[x][y] > 0 {
		return visited[x][y]
	}

	visited[x][y] = 1
	for _, dir := range dirs {
		xx, yy := x+dir[0], y+dir[1]
		visited[x][y] += dfs(grid, xx, yy, grid[x][y]) % mod
	}

	return visited[x][y]
}

/*
1
2
3
4


1, 2, 3, 4

1-2, 2-3, 3-4

1-2-3, 2-3-4

1-2-3-4
*/
