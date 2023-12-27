package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	for _, v := range board {
		for i := 0; i < len(v); i++ {
			fmt.Print(string(v[i]))
			if i != len(board)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func solveSudoku(board [][]byte) {
	solver(board, 0, 0)
}

func solver(board [][]byte, x, y int) bool {

	if x == 9 {
		return true
	}

	if y == 9 {
		return solver(board, x+1, 0)
	}
	if board[x][y] != '.' {
		return solver(board, x, y+1)
	}

	for _, v := range nums {
		if valid(board, v, x, y) {
			board[x][y] = v
			if solver(board, x, y+1) {
				return true
			}
			board[x][y] = '.'
		}
	}

	return false
}

var nums = []byte{49, 50, 51, 52, 53, 54, 55, 56, 57}

func valid(board [][]byte, num byte, x, y int) bool {

	for i := 0; i < len(board[x]); i++ {
		if board[x][i] == num {
			return false
		}
	}

	for i := 0; i < len(board); i++ {
		if board[i][y] == num {
			return false
		}
	}

	return validSquare(board, num, x, y)
}

func validSquare(board [][]byte, num byte, x, y int) bool {
	x, y = getSquare(x, y)
	for i := x * 3; i < x*3+3; i++ {
		for j := y * 3; j < y*3+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

func getSquare(x, y int) (int, int) {
	return x / 3, y / 3
}
