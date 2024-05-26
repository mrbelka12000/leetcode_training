package main

import "fmt"

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
}

func maxUniqueSplit(s string) int {
	result = 0

	runner(s, nil)

	return result
}

var result int

func runner(s string, arr []string) bool {
	if s == "" {
		result = max(result, len(arr))
		return true
	}

	for i := 0; i < len(s); i++ {
		if hasVal(arr, s[:i+1]) {
			continue
		}

		runner(s[i+1:], append(arr, s[:i+1]))
	}

	return false
}

func hasVal(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
