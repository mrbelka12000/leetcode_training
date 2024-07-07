package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSumEqual("acb", "cba", "cdb"))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var f, s, t int

	for i := 0; i < len(firstWord); i++ {
		f = f*10 + int(firstWord[i]-'a')
	}

	for i := 0; i < len(secondWord); i++ {
		s = s*10 + int(secondWord[i]-'a')
	}

	for i := 0; i < len(targetWord); i++ {
		t = t*10 + int(targetWord[i]-'a')
	}

	return f+s == t
}
