package main

import (
	"fmt"
)

func main() {
	fmt.Println(chalkReplacer([]int{3, 4, 1, 2}, 25))
}

func chalkReplacer(chalk []int, k int) int {
	var sum int

	for _, v := range chalk {
		sum += v
	}

	k = k % sum

	for i := 0; i < len(chalk); i++ {
		if chalk[i] > k {
			return i
		}
		k -= chalk[i]
	}

	return -1
}
