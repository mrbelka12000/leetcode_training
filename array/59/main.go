package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(5))
}

func generateMatrix(n int) [][]int {
	board := makeBoard(n)

	l, r := 0, n-1
	d, u := 0, n-1

	mp := make(map[[2]int]struct{})
	num := 1

	for l <= r && d <= u {
		for i := l; i <= r; i++ {
			if _, ok := mp[[2]int{d, i}]; ok {
				continue
			}
			mp[[2]int{d, i}] = struct{}{}
			board[d][i] = num
			num++
		}

		for i := d; i <= u; i++ {
			if _, ok := mp[[2]int{i, r}]; ok {
				continue
			}
			mp[[2]int{i, r}] = struct{}{}
			board[i][r] = num
			num++
		}

		for i := r; i >= l; i-- {
			if _, ok := mp[[2]int{u, i}]; ok {
				continue
			}
			mp[[2]int{u, i}] = struct{}{}
			board[u][i] = num
			num++
		}

		for i := u; i >= d; i-- {
			if _, ok := mp[[2]int{i, l}]; ok {
				continue
			}
			mp[[2]int{i, l}] = struct{}{}
			board[i][l] = num
			num++
		}

		l++
		r--
		d++
		u--
	}
	return board
}

func makeBoard(n int) [][]int {
	board := make([][]int, n)

	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	return board
}
