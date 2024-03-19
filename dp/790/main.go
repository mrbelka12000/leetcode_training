package main

import "fmt"

func main() {
	fmt.Println(numTilings(7))

}

func numTilings(n int) int {
	if n < 2 {
		return 1
	}
	i := 3
	arr := [1000]int{}

	arr[0] = 1
	arr[1] = 1
	arr[2] = 2

	for i <= n {
		arr[i] = (arr[i-1]*2 + arr[i-3]) % MOD
		i++
	}

	return arr[n]
}

const MOD = 1e9 + 7
