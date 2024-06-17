package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var res []int
	arr := make([]bool, n)

	for _, e := range edges {
		arr[e[1]] = true
	}

	for k, v := range arr {
		if !v {
			res = append(res, k)
		}
	}

	return res
}
