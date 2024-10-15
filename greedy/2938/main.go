package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSteps("101"))
	fmt.Println(minimumSteps("100"))
	fmt.Println(minimumSteps("0111"))
	fmt.Println(minimumSteps("111111111100100010"))
	fmt.Println(minimumSteps("10100000110010011010"))
}

func minimumSteps(s string) int64 {
	if len(s) == 1 {
		return 0
	}

	var (
		result int64
		n      = len(s)
	)

	var check bool
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			check = true
		}

		if s[i] == '1' {
			if check {
				result += int64(n - i - 1)
			}
			n--
		}
	}

	return result
}
