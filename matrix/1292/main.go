package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSideLength([][]int{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 4))
}

func maxSideLength(mat [][]int, ts int) int {
	var (
		result int
	)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			size := getSum(i, j, ts, mat)
			result = max(result, size)
		}
	}

	return result
}

func getSum(i, j, ts int, mat [][]int) int {

	var (
		ii, jj = i, j
	)

	for ii < len(mat) && jj < len(mat[ii]) {
		var tmpSum int

		for x := i; x <= ii; x++ {
			for y := j; y <= jj; y++ {
				tmpSum += mat[x][y]
			}
		}

		if tmpSum > ts {
			break
		}

		ii += 1
		jj += 1
	}

	return jj - j
}

/*
[[1,1,1,1]
,[1,0,0,0]
,[1,0,0,0]
,[1,0,0,0]]
*/
