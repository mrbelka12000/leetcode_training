package main

import (
	"fmt"
)

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	return runner(s1, s2, s3, make(map[[2]string]bool))
}

func runner(a, b, c string, memo map[[2]string]bool) bool {
	if c == "" {
		return a == "" && b == ""
	}

	if a == "" && b == "" {
		return false
	}

	if val, ok := memo[[2]string{a, b}]; ok {
		return val
	}
	memo[[2]string{a, b}] = false

	if a == "" {
		if b[0] == c[0] {
			return runner(a, b[1:], c[1:], memo)
		}
	} else if b == "" {
		if a[0] == c[0] {
			return runner(a[1:], b, c[1:], memo)
		}
	} else {
		if a[0] == c[0] {
			if runner(a[1:], b, c[1:], memo) {
				return true
			}
		}
		if b[0] == c[0] {
			if runner(a, b[1:], c[1:], memo) {
				return true
			}
		}
	}

	return false
}
