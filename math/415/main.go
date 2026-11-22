package main

import (
	"fmt"
)

func main() {
	fmt.Println(addStrings("456", "75"))
}

func addStrings(num1 string, num2 string) string {
	var (
		extra  int
		result string
	)

	for len(num1) > 0 && len(num2) > 0 {
		var n1, n2 int
		n1 = int(num1[len(num1)-1] - '0')
		n2 = int(num2[len(num2)-1] - '0')
		num1 = num1[:len(num1)-1]
		num2 = num2[:len(num2)-1]

		n := n1 + n2 + extra
		extra = 0
		if n >= 10 {
			n %= 10
			extra = 1
		}
		result = string(byte(n)+'0') + result
		// fmt.Println(result, n)
	}

	for i := max(len(num1), len(num2)) - 1; i >= 0; i-- {
		var n1, n2 int
		if i < len(num1) {
			n1 = int(num1[i] - '0')
		}
		if i < len(num2) {
			n2 = int(num2[i] - '0')
		}

		n := n1 + n2 + extra
		extra = 0
		if n >= 10 {
			n %= 10
			extra = 1
		}
		result = string(byte(n)+'0') + result
	}

	if extra == 1 {
		return "1" + result
	}

	return result
}
