package main

import (
	"container/heap"
	"strings"
)

func main() {

}

func topStudents(positive_feedback []string, negative_feedback []string, reports []string, student_id []int, k int) []int {
	pos := make(map[string]struct{})
	for _, v := range positive_feedback {
		pos[v] = struct{}{}
	}

	neg := make(map[string]struct{})

	for _, v := range negative_feedback {
		neg[v] = struct{}{}
	}

	pq := &PQ{}
	heap.Init(pq)

	for i, report := range reports {
		arr := strings.Split(report, " ")
		var count int
		for _, word := range arr {
			if _, ok := pos[word]; ok {
				count += 3
			}
			if _, ok := neg[word]; ok {
				count -= 1
			}
		}
		heap.Push(pq, [2]int{student_id[i], count})
	}

	result := make([]int, k)

	for i := 0; i < k; i++ {
		result[i] = heap.Pop(pq).([2]int)[0]
	}

	return result
}

type PQ [][2]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	if pq[i][1] == pq[j][1] {
		return pq[i][0] < pq[j][0]
	}
	return pq[i][1] > pq[j][1]
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.([2]int)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
