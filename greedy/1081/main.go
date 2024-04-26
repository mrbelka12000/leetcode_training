package main

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("bcabc"))
}

func smallestSubsequence(s string) string {
	freq := make(map[rune]int)

	for _, v := range s {
		freq[v]++
	}

	// fmt.Println(freq)
	present := make(map[rune]bool)
	var result []rune

	for _, c := range s {
		if !present[c] {
			for len(result) > 0 && result[len(result)-1] > c && freq[result[len(result)-1]] > 0 {
				present[result[len(result)-1]] = false
				result = result[:len(result)-1]
			}

			result = append(result, c)
			present[c] = true
		}
		freq[c]--
	}

	return string(result)
}
