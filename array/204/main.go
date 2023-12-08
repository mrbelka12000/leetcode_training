package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	arr := make([]bool, n)

	p := 2
	for p*p < n {
		inc := 2
		num := inc * p
		for num < n {
			arr[num] = true
			inc++
			num = inc * p
		}
		p++
	}

	var result int
	for _, v := range arr {
		if !v {
			result++
		}
	}

	return result - 2
}
