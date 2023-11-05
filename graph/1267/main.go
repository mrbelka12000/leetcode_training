package main

func main() {

}

func countServers(grid [][]int) int {
	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var count int
			if grid[i][j] == 1 {
				find(grid, i, j, &count)
				if count > 1 {
					result += count
				}
			}
		}
	}

	return result
}

func find(grid [][]int, x, y int, count *int) {
	grid[x][y] = 0
	*count++
	for i := 0; i < len(grid[0]); i++ {
		if grid[x][i] == 1 {
			find(grid, x, i, count)
		}
	}

	for i := 0; i < len(grid); i++ {
		if grid[i][y] == 1 {
			find(grid, i, y, count)
		}
	}

}

// [[1,0,0,1,0]
// ,[0,0,0,0,0]
// ,[0,0,0,1,0]]
