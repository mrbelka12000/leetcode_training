package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(longestDiverseString(1, 2, 3))

}

func longestDiverseString(a int, b int, c int) string {
	if a == 0 && b == 0 && c == 0 {
		return ""
	}
	pq := &PQ{}

	heap.Init(pq)

	heap.Push(pq, Info{
		B:     'a',
		Count: a,
	})

	heap.Push(pq, Info{
		B:     'b',
		Count: b,
	})

	heap.Push(pq, Info{
		B:     'c',
		Count: c,
	})

	return runner("", pq)
}

func runner(s string, pq *PQ) string {

	var arr []Info

	for pq.Len() > 0 {
		arr = append(arr, heap.Pop(pq).(Info))
	}

	if arr[0].Count == 0 {
		return s
	}

	var check bool
	// fmt.Println(arr, s)
	if len(s) >= 2 {
		if s[len(s)-1] == arr[0].B && s[len(s)-2] == arr[0].B {
			if arr[1].Count == 0 {
				return s
			}
			s += string(arr[1].B)
			arr[1].Count--
			check = true
		}
	}

	if !check {
		s += string(arr[0].B)
		arr[0].Count--
	}

	for _, v := range arr {
		heap.Push(pq, v)
	}

	return runner(s, pq)
}

type (
	PQ []Info

	Info struct {
		B     byte
		Count int
	}
)

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i].Count > pq[j].Count
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(Info)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
