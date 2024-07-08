package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheWinner(5, 2))
}

func findTheWinner(n int, k int) int {
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	tmpK := k
	for len(arr) != 1 {
		tmpK--
		val := arr[0]
		arr = arr[1:]
		if tmpK == 0 {
			tmpK = k
			continue
		}
		arr = append(arr, val)
	}

	return arr[0]
}
