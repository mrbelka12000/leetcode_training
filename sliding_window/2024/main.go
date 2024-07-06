package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxConsecutiveAnswers("TFFT", 1))
}

func maxConsecutiveAnswers(a string, k int) int {

	var (
		t, f   int
		result int
		start  int
	)

	for i := 0; i < len(a); i++ {
		if a[i] == 'T' {
			t++
		} else {
			f++
		}
		// fmt.Println(i, t, f)
		for t <= f && t > k {
			if a[start] == 'T' {
				t--
			} else {
				f--
			}
			start++
		}

		for t >= f && f > k {
			if a[start] == 'T' {
				t--
			} else {
				f--
			}
			start++
		}

		result = max(result, i-start+1)
	}

	return result
}
