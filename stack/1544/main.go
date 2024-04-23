package main

import (
	"unicode"
)

func main() {

}

func makeGood(s string) string {
	if len(s) == 1 {
		return s
	}
	var result []rune
	for _, v := range s {
		result = append(result, v)
		if len(result) > 1 && isSame(result[len(result)-1], result[len(result)-2]) {
			result = result[:len(result)-2]
		}
	}

	return string(result)
}

func isSame(r1, r2 rune) bool {
	if unicode.IsUpper(r1) {
		return unicode.ToLower(r1) == r2
	}

	if unicode.IsUpper(r2) {
		return unicode.ToLower(r2) == r1
	}

	return false
}
