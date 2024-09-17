package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

func uncommonFromSentences(s1 string, s2 string) []string {
	s := strings.Split(s1, " ")
	mp1 := make(map[string]int)
	for _, v := range s {
		mp1[v]++
	}

	s = strings.Split(s2, " ")
	mp2 := make(map[string]int)
	for _, v := range s {
		mp2[v]++
	}

	var result []string

	for k, v := range mp1 {
		if v != 1 {
			continue
		}
		if _, ok := mp2[k]; !ok {
			result = append(result, k)
		}
	}

	for k, v := range mp2 {
		if v != 1 {
			continue
		}
		if _, ok := mp1[k]; !ok {
			result = append(result, k)
		}
	}

	return result
}
