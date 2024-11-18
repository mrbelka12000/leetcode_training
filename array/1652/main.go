package main

import (
	"fmt"
)

func main() {
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}

func decrypt(code []int, k int) []int {
	n := len(code)

	result := make([]int, n)

	for i := 0; i < n; i++ {
		if k > 0 {
			tmp := k
			ind := i
			for tmp != 0 {

				ind++
				if ind >= n {
					ind = 0
				}

				result[i] += code[ind]

				tmp--
			}
		} else {
			tmp := k
			ind := i
			for tmp != 0 {
				ind--
				if ind < 0 {
					ind = n - 1
				}
				result[i] += code[ind]
				tmp++
				// fmt.Println(i, ind)
			}

		}
	}

	return result
}
