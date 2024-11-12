package main

import (
	"fmt"
)

func main() {
	fmt.Println(similarPairs([]string{"aabb", "ab", "ba"}))
}

func similarPairs(words []string) int {
	check := make(map[[26]bool]int)
	var result int

	for _, v := range words {
		var arr [26]bool

		for i := 0; i < len(v); i++ {
			arr[v[i]-'a'] = true
		}

		if val, ok := check[arr]; ok {
			result += val
		}

		check[arr]++
	}

	return result
}
