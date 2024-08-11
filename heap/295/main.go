package main

import (
	"container/heap"
	"fmt"
)

func main() {
	this := Constructor()
	this.AddNum(1)
	this.AddNum(2)
	fmt.Println(this.FindMedian())
	this.AddNum(3)
	fmt.Println(this.FindMedian())

}

type MedianFinder struct {
	smallHeap *minHeap // big heap
	bigHeap   *maxHeap // small heap
}

func Constructor() MedianFinder {
	minh := &minHeap{}
	heap.Init(minh)
	maxh := &maxHeap{}
	heap.Init(maxh)
	return MedianFinder{
		smallHeap: minh,
		bigHeap:   maxh,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.smallHeap, num)

	heap.Push(this.bigHeap, heap.Pop(this.smallHeap).(int))

	if this.smallHeap.Len() < this.bigHeap.Len() {
		heap.Push(this.smallHeap, heap.Pop(this.bigHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() (result float64) {
	l1, l2 := this.smallHeap.Len(), this.bigHeap.Len()
	if l1 == l2 {
		return float64(this.smallHeap.PQ[0]+this.bigHeap.PQ[0]) / 2
	}
	if l1 > l2 {
		return float64(this.smallHeap.PQ[0])
	}
	return float64(this.bigHeap.PQ[0])
}

type (
	minHeap struct {
		PQ
	}
	maxHeap struct {
		PQ
	}
	PQ []int
)

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq minHeap) Less(i, j int) bool {
	return pq.PQ[i] > pq.PQ[j]
}

// max-heap
func (pq maxHeap) Less(i, j int) bool {
	return pq.PQ[i] < pq.PQ[j]
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

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
