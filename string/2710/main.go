package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeTrailingZeros("51230100000"))
}

func removeTrailingZeros(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != '0' {
			return num[:i+1]
		}
	}

	return num
}
