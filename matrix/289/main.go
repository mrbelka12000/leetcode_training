package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func gameOfLife(board [][]int) {
	check := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		check[i] = make([]int, len(board[i]))
	}

	mapper(board, check, 0, 0)

	for i := 0; i < len(check); i++ {
		for j := 0; j < len(check[i]); j++ {
			if board[i][j] == 1 {
				if check[i][j] < 2 || check[i][j] > 3 {
					board[i][j] = 0
				}
			} else {
				if check[i][j] == 3 {
					board[i][j] = 1
				}
			}
		}
	}
}

/*
[[0,0,1]
,[1,0,1]
,[0,1,1]
,[1,1,1]]
*/
func mapper(board, check [][]int, x, y int) {
	if y == len(board[x]) {
		y = 0
		x++
	}
	if x == len(board) {
		return
	}

	var count int
	for _, v := range dirs {
		xx, yy := x+v[0], y+v[1]
		if xx < 0 || xx >= len(board) || yy < 0 || yy >= len(board[xx]) {
			continue
		}
		count += board[xx][yy]
	}
	check[x][y] = count
	mapper(board, check, x, y+1)
}

var dirs = [8][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
}
