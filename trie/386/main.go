package main

import "fmt"

func main() {
	fmt.Println(lexicalOrder(15))
}

func lexicalOrder(n int) []int {
	var result []int

	for i := 1; i <= 9; i++ {
		runner(n, i, &result)
	}

	return result
}

func runner(n, tmp int, arr *[]int) bool {
	if n < tmp {
		return false
	}

	*arr = append(*arr, tmp)

	for i := 0; i <= 9; i++ {
		if !runner(n, (tmp*10)+i, arr) {
			break
		}
	}

	return true
}
