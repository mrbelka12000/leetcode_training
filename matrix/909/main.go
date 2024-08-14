package main

import (
	"fmt"
)

func main() {
	fmt.Println(snakesAndLadders(make([][]int, 6)))
}

func snakesAndLadders(board [][]int) int {
	n := len(board)

	matrix := construct(n)

	q := []int{1}
	used := make([]int, n*n+1)
	var result int
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[i]
			if curr == n*n {
				return result
			}

			for i := curr + 1; i <= min(curr+6, n*n); i++ {
				if used[i] == 0 {
					used[i] = used[curr] + 1
					pos := matrix[i]
					if board[pos[0]][pos[1]] != -1 {
						q = append(q, board[pos[0]][pos[1]])
					} else {
						q = append(q, i)
					}
				}
			}
		}
		result++
		q = q[size:]
	}

	if used[n*n] == 0 {
		return -1
	}
	return used[n*n]
}

func construct(n int) map[int][2]int {
	matrix := make(map[int][2]int, n)
	num := 1
	var dir bool

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if dir {
				matrix[num] = [2]int{i, n - j - 1}
			} else {
				matrix[num] = [2]int{i, j}
			}
			num++
		}
		dir = !dir
	}

	return matrix
}

/*
 [[-1,1,2,-1]
 ,[2,13,15,-1]
 ,[-1,10,-1,-1]
 ,[-1,6,2,8]]
*/
