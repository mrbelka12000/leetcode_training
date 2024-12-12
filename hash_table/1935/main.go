package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad"))
}

func canBeTypedWords(text string, bl string) int {
	var arr [26]bool
	for i := 0; i < len(bl); i++ {
		arr[bl[i]-'a'] = true
	}

	var result int
	arrs := strings.Split(text, " ")

	for _, v := range arrs {
		var check bool
		for i := 0; i < len(v); i++ {
			if arr[v[i]-'a'] {
				check = true
				break
			}
		}
		if !check {
			result++
		}
	}

	return result
}
