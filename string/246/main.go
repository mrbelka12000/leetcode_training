package main

import (
	"fmt"
)

func main() {
	fmt.Println(isStrobogrammatic("88"))
}

func isStrobogrammatic(num string) bool {
	var tmp string

	for i := 0; i < len(num); i++ {
		if num[i] == '8' {
			tmp += "8"
		} else if num[i] == '6' {
			tmp += "9"
		} else if num[i] == '9' {
			tmp += "6"
		} else if num[i] == '1' {
			tmp += "1"
		} else if num[i] == '0' {
			tmp += "0"
		} else {
			tmp += "-"
		}
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] != num[len(num)-i-1] {
			return false
		}
	}

	return true
}
