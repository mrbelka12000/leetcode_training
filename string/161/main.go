package main

import (
	"fmt"
)

func main() {
	fmt.Println(isOneEditDistance("teacher", "detacher"))
}

func isOneEditDistance(s string, t string) bool {
	l, r := 0, 0
	for l < len(s) && r < len(t) {
		if s[l] != t[r] {
			if runner(s[l+1:], t[r:]) {
				return true
			}

			if runner(s[l+1:], t[r+1:]) {
				return true
			}

			if runner(s[l:], t[r+1:]) {
				return true
			}

			return false
		}
		l++
		r++
	}

	if len(t)-r == 1 || len(s)-l == 1 {
		return true
	}
	return false
}

func runner(s, t string) bool {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] != t[r] {
			return false
		}
		l++
		r++
	}

	return l == len(s) && r == len(t)
}
