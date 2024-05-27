package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	str := strings.Split(s, " ")

	if len(pattern) != len(str) {
		return false
	}

	used := make(map[string]bool)
	p := make(map[rune]string)

	for i, v := range pattern {
		if !used[str[i]] {
			val, ok := p[v]
			if ok && val != str[i] {
				return false
			}
			p[v] = str[i]
			used[str[i]] = true
			continue
		}

		val := p[v]
		if val != str[i] {
			return false
		}

		p[v] = str[i]

	}

	return true
}
