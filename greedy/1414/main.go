package main

import "fmt"

func main() {
	fmt.Println(findMinFibonacciNumbers(4))
}

func findMinFibonacciNumbers(k int) int {
	arr := []int{1, 1}
	n1, n2 := 0, 1
	for i := 1; i <= k; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
		if n2 > k {
			break
		}
		arr = append(arr, n2)
	}

	return runner(arr, k, 0)
}

func runner(arr []int, k, tmp int) int {
	if k == 0 {
		return tmp
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > k {
			continue
		}
		// fmt.Println(arr[i])
		return runner(arr[:i], k-arr[i], tmp+1)
	}
	return tmp
}
