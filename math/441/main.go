package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrangeCoins(5))
}

func arrangeCoins(n int) int {

	step := 1
	for n >= step {
		n -= step
		step++
	}

	return step - 1
}
