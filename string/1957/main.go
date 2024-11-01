package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

func makeFancyString(s string) string {
	if len(s) <= 2 {
		return s
	}

	str := strings.Builder{}
	str.WriteByte(s[0])
	str.WriteByte(s[1])

	for i := 2; i < len(s); i++ {
		ss := str.String()
		if ss[len(ss)-1] == s[i] && ss[len(ss)-2] == s[i] {
			continue
		}
		str.WriteByte(s[i])
	}

	return str.String()
}
