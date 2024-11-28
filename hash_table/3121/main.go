package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBCadsfka"))
}

func numberOfSpecialChars(word string) int {
	var (
		arr    [26]info
		result int
	)

	for i := 0; i < len(word); i++ {
		if word[i] <= 'Z' {
			if arr[word[i]-'A'].exists {
				arr[word[i]-'A'].last = true
				arr[word[i]-'A'].up = true
			} else {
				arr[word[i]-'A'].up = true
			}
		} else {
			if arr[word[i]-'a'].up {
				arr[word[i]-'a'].block = true
				arr[word[i]-'a'].last = false
				continue
			}
			arr[word[i]-'a'].exists = true
		}
	}

	for _, v := range arr {
		if v.last && !v.block {
			result++
		}
		// fmt.Printf("Digit: %s, info: %+v\n", string(byte(i)+'a'), v)
	}

	return result
}

type info struct {
	exists bool
	up     bool
	last   bool
	block  bool
}
