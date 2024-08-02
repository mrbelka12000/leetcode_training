package main

import (
	"fmt"
)

func main() {
	fmt.Println(appendCharacters("coaching", "coding"))
}

func appendCharacters(s string, t string) int {
	var pos int

	for i := 0; i < len(s) && pos < len(t); i++ {
		if s[i] == t[pos] {
			pos++
		}
	}

	return len(t) - pos
}
