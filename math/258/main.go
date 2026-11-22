package main

import (
	"fmt"
)

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	if num == 0 {
		return num
	}

	if num%9 == 0 {
		return 9
	}

	return num % 9
}
