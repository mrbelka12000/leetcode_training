package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(
		minimumObstacles([][]int{
			{0, 1, 1},
			{1, 1, 0},
			{1, 1, 0},
		}),
	)
}

func minimumObstacles(grid [][]int) int {
	calc := make([][]int, len(grid))
	for i := 0; i < len(calc); i++ {
		calc[i] = make([]int, len(grid[i]))
		for j := 0; j < len(calc[i]); j++ {
			calc[i][j] = math.MaxInt32
		}
	}

	calc[0][0] = 0

	pq := &PQ{}
	heap.Init(pq)
	if grid[0][0] == 1 {
		heap.Push(pq, []int{0, 0, 0, 1})
	} else {
		heap.Push(pq, []int{0, 0, 0, 0})
	}

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	result := math.MaxInt32

	for pq.Len() > 0 {

		n := heap.Pop(pq).([]int)
		if n[0] == len(grid)-1 && n[1] == len(grid[n[0]])-1 {
			return n[3]
		}

		for _, dir := range dirs {
			xx, yy := n[0]+dir[0], n[1]+dir[1]
			if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[xx]) {
				continue
			}

			var bricks, cost int
			if grid[xx][yy] == 1 {
				cost = 1
				bricks++
			}

			if calc[xx][yy] > cost+n[2] {
				calc[xx][yy] = cost + n[2]
				heap.Push(pq, []int{xx, yy, calc[xx][yy], n[3] + bricks})
			}
		}
	}

	return result
}

// x,y,cost,bricks
type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][2] < pq[j][2]
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.([]int)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
