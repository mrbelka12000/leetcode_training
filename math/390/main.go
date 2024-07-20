package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastRemaining(2134123))
}

func lastRemaining(n int) int {
	return runner(n)
}

func runner(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	val := 2 * (1 + n/2 - runner(n/2))
	return val
}
