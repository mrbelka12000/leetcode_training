package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLUSlength("aba", "caad"))
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	return max(len(a), len(b))
}
