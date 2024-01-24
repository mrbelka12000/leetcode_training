package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(matrixSum([][]int{{1, 2, 3}, {6, 5, 4}, {7, 8, 9}}))
}

type arr []int

func (a arr) Len() int {
	return len(a)
}

func (a arr) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *arr) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *arr) Pop() any {
	n := a.Len()
	tmp := (*a)[n-1]
	*a = (*a)[:n-1]
	return tmp
}

func matrixSum(nums [][]int) int {
	var heaps []*arr
	for i := 0; i < len(nums); i++ {
		a := &arr{}
		for j := 0; j < len(nums[i]); j++ {
			heap.Push(a, nums[i][j])
		}
		heaps = append(heaps, a)
	}

	var result int

	for heaps[0].Len() != 0 {
		var max int

		for _, v := range heaps {
			//fmt.Println(*v)
			val := heap.Pop(v).(int)
			if max < val {
				max = val
			}
		}
		//fmt.Println()
		result += max
	}

	return result
}
