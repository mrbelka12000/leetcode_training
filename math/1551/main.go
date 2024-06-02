package main

import "fmt"

func main() {
	fmt.Println(minOperations(6))
}

func minOperations(n int) int {
	if n == 1 {
		return 0
	}

	arr := make([]int, 0, n)
	var sum int
	for i := 0; i < n; i++ {
		val := (2 * i) + 1
		arr = append(arr, val)
		sum += val
	}

	target := sum / n

	var result int

	for i := 0; i < len(arr)/2; i++ {
		result += abs(arr[i] - target)
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
