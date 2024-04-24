package main

import "container/heap"

func main() {

}

func leastInterval(tasks []byte, k int) int {
	freq := make(map[byte]int)

	for _, v := range tasks {
		freq[v]++
	}

	pq := &PQ{}

	for k, v := range freq {
		heap.Push(pq, task{
			val:   k,
			count: v,
		})
	}

	var result int
	for pq.Len() > 0 {
		var works []task
		n := k + 1

		for n > 0 && pq.Len() > 0 {
			val := heap.Pop(pq).(task)
			result++

			if val.count > 1 {
				val.count--
				works = append(works, val)
			}
			n--
		}

		for _, v := range works {
			heap.Push(pq, v)
		}

		if pq.Len() > 0 {
			result += n
		}
	}

	return result
}

type (
	task struct {
		val   byte
		count int
	}
)

type PQ []task

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(task)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
