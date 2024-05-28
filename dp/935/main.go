package main

import "fmt"

func main() {
	fmt.Println(knightDialer(3131))
}

func knightDialer(n int) int {
	var sum int
	mp = make(map[[2]int]int)

	for i := 0; i <= 9; i++ {
		sum += runner(n, i)
	}

	return sum % mod
}

var (
	mp  map[[2]int]int
	mod = 1_000_000_007
)

func runner(n, num int) int {
	if n == 1 {
		return 1
	}

	if val, ok := mp[[2]int{n, num}]; ok {
		return val
	}

	if num == 0 {
		mp[[2]int{n, num}] = (runner(n-1, 4) + runner(n-1, 6)%mod)
	} else if num == 1 {
		mp[[2]int{n, num}] = (runner(n-1, 8) + runner(n-1, 6)%mod)
	} else if num == 2 {
		mp[[2]int{n, num}] = (runner(n-1, 7) + runner(n-1, 9)%mod)
	} else if num == 3 {
		mp[[2]int{n, num}] = (runner(n-1, 4) + runner(n-1, 8)%mod)
	} else if num == 4 {
		mp[[2]int{n, num}] = (runner(n-1, 9) + runner(n-1, 3) + runner(n-1, 0)%mod)
	} else if num == 6 {
		mp[[2]int{n, num}] = (runner(n-1, 1) + runner(n-1, 7) + runner(n-1, 0)%mod)
	} else if num == 7 {
		mp[[2]int{n, num}] = (runner(n-1, 6) + runner(n-1, 2)%mod)
	} else if num == 8 {
		mp[[2]int{n, num}] = (runner(n-1, 1) + runner(n-1, 3)%mod)
	} else if num == 9 {
		mp[[2]int{n, num}] = (runner(n-1, 4) + runner(n-1, 2)%mod)
	}

	return mp[[2]int{n, num}] % int(mod)
}
