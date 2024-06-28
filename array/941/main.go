package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	var count int

	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == arr[i] || arr[i+1] == arr[i] {
			return false
		}

		if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
			count++
		}

		if arr[i-1] > arr[i] && arr[i+1] > arr[i] {
			return false
		}
	}

	return count == 1
}
