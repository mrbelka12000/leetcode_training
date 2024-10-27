package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

func mostCommonWord(p string, banned []string) string {
	p = strings.ToLower(p)
	mp := make(map[string]bool)
	for _, v := range banned {
		mp[v] = true
	}

	p = strings.Replace(p, ".", " ", -1)
	p = strings.Replace(p, ",", " ", -1)
	p = strings.Replace(p, "?", " ", -1)
	p = strings.Replace(p, ";", " ", -1)
	p = strings.Replace(p, "'", " ", -1)
	p = strings.Replace(p, "!", " ", -1)

	count := make(map[string]int)
	arr := strings.Split(p, " ")
	for _, v := range arr {
		if v == "" {
			continue
		}
		count[v]++
	}

	for k := range mp {
		delete(count, k)
	}

	var (
		result string
		rCount int
	)

	for k, v := range count {
		if v > rCount {
			rCount = v
			result = k
		}
	}

	return result
}
