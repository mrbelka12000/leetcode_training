package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func maxSatisfaction(sat []int) int {
	sort.Ints(sat)

	var result int

	for i := 0; i < len(sat); i++ {
		result = max(result, summer(sat[i:]))
	}

	return result
}

func summer(nums []int) int {
	var sum int
	mult := 1

	for _, v := range nums {
		sum += (v * mult)
		mult++
	}

	return sum
}
