package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("((("))

}

func minAddToMakeValid(s string) int {

	var countO, result int

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			countO++
		case ')':
			if countO == 0 {
				result++
			} else {
				countO--
			}
		}
	}

	if countO != 0 {
		result += countO
	}

	return result
}
