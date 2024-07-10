package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
}

func minOperations(logs []string) int {
	var stack []string

	for _, v := range logs {
		if v[0] == '.' && v[1] == '.' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		if v[0] != '.' {
			stack = append(stack, v)
		}
	}

	return len(stack)
}
