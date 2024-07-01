package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeConsecutiveOdds([]int{2, 6, 4, 1}))
}

func threeConsecutiveOdds(arr []int) bool {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1]%2 == 1 && arr[i]%2 == 1 && arr[i+1]%2 == 1 {
			return true
		}
	}

	return false
}
