package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func removeAnagrams(words []string) []string {
	for {

		prev := [26]int{}
		for _, v := range words[0] {
			prev[v-'a']++
		}

		var check bool
		for i := 1; i < len(words); i++ {
			curr := [26]int{}
			for _, v := range words[i] {
				curr[v-'a']++
			}

			if curr == prev {
				words = remove(words, i)
				check = true
				i--
				continue
			}
			prev = curr
		}

		if !check {
			break
		}

	}

	return words
}

func remove(arr []string, i int) []string {
	return append(arr[:i], arr[i+1:]...)
}
