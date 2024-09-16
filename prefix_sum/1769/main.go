package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations("001011"))
}

func minOperations(boxes string) []int {
	n := len(boxes)
	var count, total int

	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = total
		count += int(boxes[i] - '0')
		total += count
	}

	count, total = 0, 0
	for i := n - 1; i >= 0; i-- {
		result[i] += total
		count += int(boxes[i] - '0')
		total += count
	}
	return result
}
