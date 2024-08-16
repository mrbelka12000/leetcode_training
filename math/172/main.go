package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trailingZeroes(1000))
}

func trailingZeroes(n int) int {
	var result int
	del := 5.0
	pow := 1.0

	div := math.MaxInt32
	for div > 1 {
		div = n / int(math.Pow(del, pow))
		result += div
		pow++
	}

	return result
}
