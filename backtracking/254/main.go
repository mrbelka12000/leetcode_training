package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func getFactors(n int) [][]int {
	return runner(2, n, nil)
}

func runner(pos, n int, tmp []int) [][]int {
	// fmt.Println(n, tmp)
	if n == 1 {
		if len(tmp) <= 1 {
			return nil
		}
		cp := make([]int, len(tmp))
		copy(cp, tmp)
		return [][]int{cp}
	}

	result := make([][]int, 0, 10)
	for i := pos; i <= n; i++ {
		if n%i == 0 {
			val := runner(i, n/i, append(tmp, i))
			if val != nil {
				result = append(result, val...)
			}
		}
	}

	return result
}
