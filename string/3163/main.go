package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))

}

func compressedString(word string) string {
	var (
		result strings.Builder
		count  = 1
	)

	for i := 1; i < len(word); i++ {
		if word[i] != word[i-1] || count == 9 {
			result.WriteString(fmt.Sprintf("%v%v", count, string(word[i-1])))
			count = 1
			continue
		}
		count++
	}

	result.WriteString(fmt.Sprintf("%v%v", count, string(word[len(word)-1])))

	return result.String()
}
