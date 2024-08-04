package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func minFlips(mat [][]int) int {
	if len(mat) == 1 && len(mat[0]) == 1 {
		return mat[0][0]
	}
	result = math.MaxInt32

	var count int
	visited := make([][]bool, len(mat))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			count += mat[i][j]
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			visited[i][j] = true
			runner(mat, visited, i, j, 0, count)
			visited[i][j] = false
		}
	}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func runner(mat [][]int, visited [][]bool, x, y, steps, count int) {
	if count == 0 {
		result = min(result, steps)
		return
	}

	if mat[x][y] == 1 {
		count--
	} else {
		count++
	}
	mat[x][y] = unempty(mat[x][y])
	for _, v := range dirs {
		xx, yy := v[0]+x, v[1]+y
		if xx < 0 || xx >= len(mat) || yy < 0 || yy >= len(mat[xx]) {
			continue
		}
		if mat[xx][yy] == 1 {
			count--
		} else {
			count++
		}

		mat[xx][yy] = unempty(mat[xx][yy])
	}

	if count == 0 {
		result = min(result, steps+1)
	}

	visited[x][y] = true

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if !visited[i][j] {
				runner(mat, visited, i, j, steps+1, count)
			}
		}
	}

	for _, v := range dirs {
		xx, yy := v[0]+x, v[1]+y
		if xx < 0 || xx >= len(mat) || yy < 0 || yy >= len(mat[xx]) {
			continue
		}
		if mat[xx][yy] == 1 {
			count++
		} else {
			count--
		}

		mat[xx][yy] = unempty(mat[xx][yy])
	}

	mat[x][y] = unempty(mat[x][y])
	count += mat[x][y]
	visited[x][y] = false
}

var result int

func unempty(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}
