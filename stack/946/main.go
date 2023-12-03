package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	var st []int

	for _, v := range pushed {
		for i := len(st) - 1; i >= 0; i-- {
			if st[i] == popped[0] {
				popped = popped[1:]
				st = st[:len(st)-1]
			}
		}
		st = append(st, v)
	}

	for i := len(st) - 1; i >= 0; i-- {
		if st[i] == popped[0] {
			popped = popped[1:]
			st = st[:len(st)-1]
		}
	}

	return len(st) == 0
}
