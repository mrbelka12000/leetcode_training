package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func minTimeToReach(moveTime [][]int) int {
	grid := make([][]info, len(moveTime))
	for i := 0; i < len(moveTime); i++ {
		grid[i] = make([]info, len(moveTime[i]))
	}
	n := len(moveTime)
	m := len(moveTime[0])

	q := []info{
		{
			0, 0, 0,
		},
	}
	calc := make([][]int, n)
	for i := 0; i < n; i++ {
		calc[i] = make([]int, m)
		for j := 0; j < m; j++ {
			calc[i][j] = math.MaxInt32
		}
	}

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) != 0 {

		curr := q[0]
		q = q[1:]
		for _, dir := range dirs {
			xx, yy, cost := curr.x+dir[0], curr.y+dir[1], curr.t
			if xx < 0 || xx >= n || yy < 0 || yy >= m {
				continue
			}

			if calc[xx][yy] > cost+1 {
				nCost := cost + 1
				if moveTime[xx][yy] > cost {
					nCost = moveTime[xx][yy] + 1
				}
				calc[xx][yy] = nCost
				q = append(q, info{xx, yy, nCost})
			}
		}
	}

	return calc[n-1][m-1]
}

type info struct {
	x, y, t int
}
