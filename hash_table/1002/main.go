package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(words []string) []string {
	var arr [][26]int

	for _, v := range words[1:] {
		var tmp [26]int
		for i := 0; i < len(v); i++ {
			tmp[v[i]-'a']++
		}
		arr = append(arr, tmp)
	}

	var result []string
	for i := 0; i < len(words[0]); i++ {
		count := 1

		for j := 0; j < len(arr); j++ {
			if arr[j][words[0][i]-'a'] > 0 {
				arr[j][words[0][i]-'a']--
				count++
			}
		}

		if count == len(words) {
			result = append(result, string(words[0][i]))
		}
	}

	return result
}
