package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3}))
}

func rearrangeBarcodes(barcodes []int) []int {

	mp := make(map[int]int)
	for _, v := range barcodes {
		mp[v]++
	}

	pq := &PQ{}
	heap.Init(pq)
	for k, v := range mp {
		heap.Push(pq, []int{k, v})
	}

	return runner(pq, len(barcodes))
}

func runner(pq *PQ, size int) []int {
	result := make([]int, 0, size)
	for pq.Len() != 0 {
		val := heap.Pop(pq).([]int)
		if len(result) == 0 || result[len(result)-1] != val[0] {
			val[1]--
			result = append(result, val[0])
			if val[1] > 0 {
				heap.Push(pq, val)
			}
			continue
		} else {
			second := heap.Pop(pq).([]int)
			second[1]--
			result = append(result, second[0])
			if second[1] > 0 {
				heap.Push(pq, second)
			}
		}

		heap.Push(pq, val)
	}

	return result
}

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i][1] > pq[j][1]
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
