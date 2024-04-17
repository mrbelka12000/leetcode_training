package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3))
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	var (
		arr    [][2]int
		result int64
	)

	for i := 0; i < len(nums1); i++ {
		arr = append(arr, [2]int{
			nums1[i],
			nums2[i],
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	p := &PQ{}

	heap.Init(p)
	var sum int

	for i := 0; i < k; i++ {
		sum += arr[i][0]
		heap.Push(p, arr[i])
	}

	result = int64(sum * arr[k-1][1])

	for i := k; i < len(arr); i++ {
		smallNum := heap.Pop(p).([2]int)
		sum += arr[i][0] - smallNum[0]

		heap.Push(p, arr[i])

		result = max(result, int64(sum*arr[i][1]))
	}

	fmt.Println(p)

	return result
}

type PQ [][2]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
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
