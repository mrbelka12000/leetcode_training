package main

import (
	"fmt"
)

func main() {
	fmt.Println(3)
}

func countArrangement(n int) int {
	var runner func(arr []int, used []bool, pos int, path []int) int

	runner = func(arr []int, used []bool, pos int, path []int) (result int) {
		if pos == len(used) {
			return 1
		} else {
			for _, v := range arr {
				if !used[v] && (v%pos == 0 || pos%v == 0) {
					used[v] = true
					result += runner(arr, used, pos+1, append(path, v))
					used[v] = false
				}
			}
		}
		return result
	}

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	return runner(arr, make([]bool, n+1), 1, nil)
}
