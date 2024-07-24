package main

import (
	"fmt"
)

func main() {
	fmt.Println(selfDividingNumbers(12, 55))
}

func selfDividingNumbers(left int, right int) []int {
	var result []int

	for i := left; i <= right; i++ {
		if is(i) {
			result = append(result, i)
		}
	}

	return result
}

func is(num int) bool {
	numStr := fmt.Sprint(num)
	for i := 0; i < len(numStr); i++ {
		val := int(numStr[i] - '0')
		if val == 0 {
			return false
		}
		if num%val != 0 {
			return false
		}
	}
	return true
}
