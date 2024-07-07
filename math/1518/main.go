package main

import (
	"fmt"
)

func main() {
	fmt.Println(numWaterBottles(15, 4))
}

func numWaterBottles(numBottles int, numExchange int) int {
	var empty, result int
	for numBottles > 0 {
		result += numBottles
		n := numBottles + empty
		numBottles = n / numExchange
		empty = n % numExchange
	}

	return result
}
