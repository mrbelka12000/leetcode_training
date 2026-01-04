package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumFourDivisors([]int{21, 4, 7}))
}

func sumFourDivisors(nums []int) int {
	var result int

	for _, v := range nums {
		var (
			tmp int
		)

		tmp += 1 + v

		mp := make(map[int]struct{})
		for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
			if v%i == 0 {
				tmp += i
				mp[i] = struct{}{}
				_, ok := mp[v/i]
				if !ok {
					tmp += v / i
				}
				mp[v/i] = struct{}{}
				if len(mp) > 2 {
					break
				}
			}
		}
		if len(mp) == 2 {
			result += tmp
		}
	}

	return result
}
