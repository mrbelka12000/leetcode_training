package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var result [][]string
	mp = make(map[[2]int]struct{})

	num := n
	board := makeBoard(n)
	solver(board, n, 0, 0, &num, &result)

	return result
}

var mp = make(map[[2]int]struct{})

func castToString(board [][]byte) []string {
	arr := make([]string, len(board))

	for i := 0; i < len(board); i++ {
		arr[i] = string(board[i])
	}

	return arr
}

func solver(board [][]byte, n, x, y int, num *int, result *[][]string) bool {
	if *num == 0 {
		tmp := make([][]byte, n)

		copy(tmp, board)

		*result = append(*result, castToString(tmp))
		return false
	}

	for i := x; i < n; i++ {
		for j := y; j < n; j++ {
			if _, ok := mp[[2]int{i, j}]; ok {
				continue
			}

			if isValid(board, i, j) {
				board[i][j] = 'Q'
				mp[[2]int{i, j}] = struct{}{}
				*num--
				if solver(board, n, i+1, y, num, result) {
					return true
				}
				delete(mp, [2]int{i, j})
				board[i][j] = '.'
				*num++
			}
		}
	}

	return false
}

func makeBoard(n int) [][]byte {
	board := make([][]byte, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i] = append(board[i], '.')
		}
	}

	return board
}

func isValid(board [][]byte, x, y int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'Q' {
				if i == x {
					return false
				}
				if j == y {
					return false
				}

				if abs(i-x) == abs(j-y) {
					return false
				}
			}
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
