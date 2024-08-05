package main

import (
	"fmt"
)

func main() {
	fmt.Println(selfDivisiblePermutationCount(5))
}

func selfDivisiblePermutationCount(n int) int {
	nums = nil
	for i := 0; i <= n; i++ {
		nums = append(nums, i)
	}

	result = 0
	runner(nums, make([]bool, n+1), 1, nil)

	return result
}

func runner(arr []int, used []bool, n int, tmp []int) {
	if n == len(used) {
		result++
	} else {
		for i, v := range nums {
			if i == 0 {
				continue
			}
			if !used[v] && gcd(v, n) == 1 {
				used[v] = true
				runner(append(arr, v), used, n+1, append(tmp, v))
				used[v] = false
			}
		}
	}
}

var (
	result int
	nums   []int
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
