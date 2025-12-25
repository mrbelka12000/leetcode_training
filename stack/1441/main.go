package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
}

func buildArray(target []int, n int) []string {
	var result []string
	for i := 1; i <= n && len(target) > 0; i++ {
		result = append(result, pu)
		if target[0] != i {
			result = append(result, pp)
			continue
		}
		target = target[1:]
	}

	return result
}

const (
	pu = "Push"
	pp = "Pop"
)
