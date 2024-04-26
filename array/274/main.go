package main

import "fmt"

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(cit []int) int {
	arr := make([]int, len(cit)+1)

	for _, v := range cit {
		if v >= len(cit) {
			arr[len(cit)]++
		} else {
			arr[v]++
		}
	}

	pos := 0
	for i, v := range arr {
		for j := 0; j < v; j++ {
			cit[pos] = i
			pos++
		}
	}

	hIndex := 1
	for i := len(cit) - 1; i >= 0; i-- {
		if cit[i] < hIndex {
			return hIndex - 1
		}
		hIndex++
	}

	return len(cit)
}
