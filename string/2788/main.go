package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(splitWordsBySeparator([]string{"$easy$", "$problem$"}, '$'))
}

func splitWordsBySeparator(words []string, separator byte) []string {
	var result []string

	for _, v := range words {
		arr := strings.Split(v, string(separator))

		for _, word := range arr {
			if word != "" {
				result = append(result, word)
			}
		}
	}

	return result
}
