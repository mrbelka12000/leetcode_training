package main

import "fmt"

func main() {
	//fmt.Println(longestIncreasingPath([][]int{
	//	{9, 9, 4},
	//	{6, 6, 8},
	//	{2, 1, 1}}))
	fmt.Println(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1}}))
}

func longestIncreasingPath(matrix [][]int) int {
	var runner func(x, y, d int)
	memo := make(map[[2]int]int)
	runner = func(x, y, d int) {
		if val, ok := memo[[2]int{x, y}]; ok {
			if val < d {
				memo[[2]int{x, y}] = d
			} else {
				return
			}
		} else {
			memo[[2]int{x, y}] = d
		}
		for _, dir := range dirs {
			xx, yy := x+dir[0], y+dir[1]
			if xx < 0 || xx >= len(matrix) || yy < 0 || yy >= len(matrix[xx]) || matrix[x][y] >= matrix[xx][yy] {
				continue
			}
			runner(xx, yy, d+1)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			runner(i, j, 1)
			//fmt.Println(memo)
			//return 0
		}
	}

	var result int
	for _, v := range memo {
		if result < v {
			result = v
		}
	}

	//fmt.Println(matrix)
	return result
}

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
