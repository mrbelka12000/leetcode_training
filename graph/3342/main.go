package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minTimeToReach())
}

func minTimeToReach(moveTime [][]int) int {
	grid := make([][]info, len(moveTime))
	for i := 0; i < len(moveTime); i++ {
		grid[i] = make([]info, len(moveTime[i]))
	}
	n := len(moveTime)
	m := len(moveTime[0])

	pq := &PQ{}
	heap.Init(pq)

	heap.Push(pq, info{0, 0, 0, 0, 1})
	calc := make([][]int, n)
	for i := 0; i < n; i++ {
		calc[i] = make([]int, m)
		for j := 0; j < m; j++ {
			calc[i][j] = math.MaxInt32
		}
	}

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for pq.Len() != 0 {

		curr := heap.Pop(pq).(info)

		if curr.x == n-1 && curr.y == m-1 {
			break
		}

		for _, dir := range dirs {
			xx, yy, cost, temp, step := curr.x+dir[0], curr.y+dir[1], curr.t, curr.cost, curr.step
			if xx < 0 || xx >= n || yy < 0 || yy >= m {
				continue
			}

			if calc[xx][yy] > cost+1 {
				nStep := step + 1
				if nStep == 2 {
					temp = 1
					nStep = 0
				}
				nCost := cost + temp
				if moveTime[xx][yy] > cost {
					nCost = moveTime[xx][yy] + temp
				}
				calc[xx][yy] = nCost
				heap.Push(pq, info{xx, yy, nCost, temp + 1, nStep})
			}
		}
		// fmt.Println(pq)
	}

	return calc[n-1][m-1]
}

type info struct {
	x, y, t, cost, step int
}

type PQ []info

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i].t < pq[j].t
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(info)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}

/*
[[0,0,0,0]
,[0,0,0,0]]

*/
