package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PQ{}
	var result [][]int
	heap.Init(pq)
	visited := make(map[[2]int]bool)
	heap.Push(pq, []int{nums1[0] + nums2[0], 0, 0})
	visited[[2]int{0, 0}] = true

	for k > 0 && pq.Len() > 0 {
		val := heap.Pop(pq).([]int)

		i, j := val[1], val[2]
		result = append(result, []int{nums1[i], nums2[j]})

		if i+1 < len(nums1) && !visited[[2]int{i + 1, j}] {
			visited[[2]int{i + 1, j}] = true
			heap.Push(pq, []int{nums1[i+1] + nums2[j], i + 1, j})
		}

		if j+1 < len(nums2) && !visited[[2]int{i, j + 1}] {
			visited[[2]int{i, j + 1}] = true
			heap.Push(pq, []int{nums1[i] + nums2[j+1], i, j + 1})
		}
		k--
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
	return pq[i][0] < pq[j][0]
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
