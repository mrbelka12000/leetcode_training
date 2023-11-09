package main

import "fmt"

func main() {
	fmt.Println(countVowelStrings(5))
}

func countVowelStrings(n int) int {
	result := 1
	for i := 1; i < 5; i++ {
		result *= n + i
	}

	// fmt.Println(result)

	return result / 24
}

//this sum of A[n] + E[n} ... + U[n] simplifies into (n+1)(n+2)(n+3)*(n+4) / 24.
