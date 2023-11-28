package main

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
}

func numPairsDivisibleBy60(time []int) int {
	var result int

	for i := 0; i < len(time); i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				result++
			}
		}
	}

	return result
}
