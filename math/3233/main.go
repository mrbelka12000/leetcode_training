package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func nonSpecialCount(l int, r int) int {
	if len(primes) == 0 {
		getPrimes()
	}

	result := r - l + 1
	for _, v := range primes {
		num := v * v
		if num >= l && num <= r {
			result--
		}
	}

	return result
}

var primes []int

func getPrimes() {

	n := int(math.Sqrt(1e9))
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

	for i, v := range arr[2:] {
		if !v {
			primes = append(primes, i+2)
		}
	}

	return
}
