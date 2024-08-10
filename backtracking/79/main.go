package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if runner(board, 1, i, j, word, used) {
					return true
				}
			}
		}
	}

	return false
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func runner(board [][]byte, ind, x, y int, word string, used [][]bool) bool {
	if ind == len(word) {
		return true
	}

	used[x][y] = true

	for _, v := range dirs {
		xx, yy := x+v[0], y+v[1]

		if xx < 0 || xx >= len(board) || yy < 0 || yy >= len(board[xx]) || used[xx][yy] {
			continue
		}

		if board[xx][yy] != word[ind] {
			continue
		}

		if runner(board, ind+1, xx, yy, word, used) {
			return true
		}
	}

	used[x][y] = false

	return false
}
