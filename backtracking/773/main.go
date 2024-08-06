package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func slidingPuzzle(board [][]int) int {
	result = math.MaxInt32

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				runner(board, i, j, 0, -1)
				if result == math.MaxInt32 {
					return -1
				}
				return result
			}
		}
	}

	return -1
}

func runner(board [][]int, x, y, cost, prev int) {
	if cost > 18 {
		return
	}
	if isValid(board) {
		result = min(result, cost)
		return
	}

	for i, v := range dirs {
		if i == prev {
			continue
		}
		xx, yy := x+v[0], y+v[1]
		if xx < 0 || xx >= len(board) || yy < 0 || yy >= len(board[xx]) {
			continue
		}

		board[xx][yy], board[x][y] = board[x][y], board[xx][yy]
		runner(board, xx, yy, cost+1, getPrev(i))
		board[xx][yy], board[x][y] = board[x][y], board[xx][yy]
	}
}

func getPrev(i int) int {
	switch i {
	case 0:
		return 1
	case 1:
		return 0
	case 2:
		return 3
	default:
		return 2
	}
}

var (
	result int
	dirs   = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func isValid(board [][]int) bool {
	return board[0][0] == 1 && board[0][1] == 2 && board[0][2] == 3 && board[1][0] == 4 && board[1][1] == 5
}
