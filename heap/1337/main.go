package main

import "container/heap"

func main() {

}

func kWeakestRows(mat [][]int, k int) []int {
	pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < len(mat); i++ {
		var count int
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				count++
			}
		}
		heap.Push(pq, []int{count, i})
	}

	result := make([]int, k)

	for i := 0; i < k; i++ {
		// fmt.Println(pq)
		result[i] = heap.Pop(pq).([]int)[1]
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

// min-heap
func (pq PQ) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] < pq[j][1]
	}
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
