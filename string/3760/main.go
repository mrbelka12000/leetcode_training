package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistinct("ababa"))
}

func maxDistinct(s string) int {
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	var result int

	for _, v := range arr {
		if v > 0 {
			result++
		}
	}

	return result
}
