package main

import "container/heap"

func main() {

}

func getKth(lo int, hi int, k int) int {
	memo = make(map[int]int)

	// var sl [][2]int

	// for i := hi ; i >= lo ;i--{
	//     sl = append(sl, [2]int{getPower(i),i})
	// }

	// sort.Slice(sl, func(i, j int)bool{
	//     if sl[i][0] == sl[j][0]{
	//         return sl[i][1] < sl[j][1]
	//     }
	//     return sl[i][0] < sl[j][0]
	// })

	pq := &PQ{}
	heap.Init(pq)

	for i := hi; i >= lo; i-- {
		heap.Push(pq, []int{getPower(i), i})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	// fmt.Println(pq)

	// for i := 0 ; i < k-1;i++{
	// heap.Pop(pq)
	// }

	return (*pq)[0][1]
	// return sl[k-1][1]
}

var memo map[int]int

func getPower(x int) int {
	if x == 1 {
		return 0
	}
	if val, ok := memo[x]; ok {
		return val
	}

	memo[x] = 1
	if x%2 == 0 {
		memo[x] += getPower(x / 2)
	} else {
		memo[x] += getPower(3*x + 1)
	}

	return memo[x]
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
		return pq[i][1] > pq[j][1]
	}
	return pq[i][0] > pq[j][0]
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
