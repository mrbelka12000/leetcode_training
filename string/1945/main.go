package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLucky("leetcode", 2))
}

func getLucky(s string, k int) int {
	var (
		result = getSumStr(s)
	)
	k--
	for k > 0 {
		k--
		result = getSumInt(result)
	}

	return result
}

func getSumInt(i int) int {
	var tmp int
	for i > 0 {
		tmp += i % 10
		i /= 10
	}
	return tmp
}
func getSumStr(s string) int {
	var tmp int
	for i := 0; i < len(s); i++ {
		tmp += (int(s[i]) - 'a' + 1) / 10
		tmp += (int(s[i]) - 'a' + 1) % 10
	}
	return tmp
}
