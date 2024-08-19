package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	mp := make(map[int]int)
	for _, v := range hand {
		mp[v]++
	}
	pq := &PQ{}
	heap.Init(pq)
	for k, v := range mp {
		heap.Push(pq, []int{v, k})
	}

	for pq.Len() > 0 {
		var (
			tmp  [][]int
			prev = -1
			i    int
		)

		for i = 0; i < groupSize && pq.Len() > 0; i++ {
			val := heap.Pop(pq).([]int)

			if prev == -1 || val[1]-prev == 1 {
				prev = val[1]
			} else {
				return false
			}
			val[0]--
			if val[0] > 0 {
				tmp = append(tmp, []int{val[0], val[1]})
			}
		}

		if i != groupSize {
			return false
		}

		for _, v := range tmp {
			heap.Push(pq, v)
		}
	}
	return true
}

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i][1] < pq[j][1]
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
