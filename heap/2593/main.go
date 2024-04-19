package main

import "container/heap"

func main() {

}

func findScore(nums []int) int64 {
	pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < len(nums); i++ {
		heap.Push(pq, [2]int{nums[i], i})
	}

	arr := make([]bool, len(nums))
	var result int64
	for pq.Len() > 0 {
		val := heap.Pop(pq).([2]int)

		if arr[val[1]] {
			continue
		}

		result += int64(val[0])
		arr[val[1]] = true
		if val[1] != 0 {
			arr[val[1]-1] = true
		}
		if val[1] != len(arr)-1 {
			arr[val[1]+1] = true
		}
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

// min-heap
func (pq PQ) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] < pq[j][1]
	}
	return pq[i][0] < pq[j][0]
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
