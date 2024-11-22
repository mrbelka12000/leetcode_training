package main

import (
	"fmt"
)

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	var arr [21]int

	for _, bill := range bills {
		diff := bill - 5
		for diff != 0 {
			for diff >= 20 && arr[20] > 0 {
				diff -= 20
				arr[20]--
			}

			for diff >= 10 && arr[10] > 0 {
				diff -= 10
				arr[10]--
			}

			for diff >= 5 && arr[5] > 0 {
				diff -= 5
				arr[5]--
			}

			if diff != 0 {
				return false
			}
		}

		arr[bill]++
	}

	return true
}
