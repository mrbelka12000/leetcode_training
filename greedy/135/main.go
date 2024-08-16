package main

import (
	"fmt"
)

func main() {
	fmt.Println(candy([]int{1, 6, 10, 8, 7, 3, 2}))
}

func candy(ratings []int) int {
	n := len(ratings)
	cn := make([]int, n)

	for i := 0; i < n-1; i++ {
		if ratings[i+1] > ratings[i] {
			cn[i+1] = cn[i] + 1
		}
	}

	for i := n - 1; i >= 1; i-- {
		if ratings[i-1] > ratings[i] && cn[i-1] <= cn[i] {
			cn[i-1] = cn[i] + 1
		}
	}

	var result int
	for _, v := range cn {
		result += v
	}
	return result + n
}
