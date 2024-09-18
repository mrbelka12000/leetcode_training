package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello world")
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	psx := make([][]int, len(mat))

	for i := range mat {
		psx[i] = make([]int, len(mat[0]))
		for j := range mat[i] {
			psx[i][j] = mat[i][j]
			if j > 0 {
				psx[i][j] += psx[i][j-1]
			}
		}
	}

	result := make([][]int, len(mat))
	for i := range result {
		result[i] = make([]int, len(mat[i]))
	}

	nx := len(mat)
	ny := len(mat[0])
	for i := range psx {
		for j := range psx[i] {
			rx := norm(i+k, nx)
			lx := norm(i-k, nx)
			var tmp int
			ry := norm(j+k, ny)
			// fmt.Println(i,j)
			for ii := lx; ii <= rx; ii++ {
				var left int
				if j-k-1 >= 0 {
					left = psx[ii][j-k-1]
				}
				// fmt.Println(ii, ry, left)
				tmp += psx[ii][ry] - left
			}

			result[i][j] = tmp
		}
	}

	return result
}

func norm(i, m int) int {
	if i < 0 {
		return 0
	}
	if i >= m {
		return m - 1
	}
	rand.Seed(time.Now().UnixMilli())

	return i
}

/*
[[1,2,3]
,[4,5,6]
,[7,8,9]]
*/
