package main

import (
	"fmt"
)

func main() {
	fmt.Println(scoreOfString("Hello World"))
}

func scoreOfString(s string) (result int) {
	for i := 0; i < len(s)-1; i++ {
		result += abs(int(s[i]) - int(s[i+1]))
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
