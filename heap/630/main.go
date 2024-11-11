package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	pq := &PQ{}
	heap.Init(pq)
	var (
		result int
		t      int
	)
	// fmt.Println(courses)
	for i := 0; i < len(courses); i++ {
		curr := courses[i]
		if curr[0] > curr[1] {
			continue
		}
		if t+curr[0] > curr[1] {
			last := heap.Pop(pq).(int)
			if last < curr[0] {
				heap.Push(pq, last)
				continue
			}
			t -= last
			result--
		}

		t += curr[0]
		result++
		heap.Push(pq, curr[0])
		// fmt.Println(t)
	}

	return result
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(int)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
