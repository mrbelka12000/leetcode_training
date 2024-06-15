package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu"))
}

func longestBeautifulSubstring(word string) int {
	var (
		result   int
		lastChar byte
		count    int
		arr      []byte
	)

	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			if word[i] < lastChar {
				arr = nil
				count = 0
			}

			lastChar = word[i]
			count++
			if len(arr) == 0 || arr[len(arr)-1] != word[i] {
				arr = append(arr, word[i])
			}
			if len(arr) == 5 {
				result = max(result, count)
			}
		}
	}

	return result
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}
