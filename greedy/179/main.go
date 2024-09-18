package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

func largestNumber(nums []int) string {
	var numStr []string
	for _, v := range nums {
		numStr = append(numStr, fmt.Sprint(v))
	}

	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i]+numStr[j] > numStr[j]+numStr[i]
	})
	if numStr[0] == "0" {
		return "0"
	}

	return strings.Join(numStr, "")
}
