package main

import "fmt"

func main() {
	atm := Constructor()

	atm.Deposit([]int{11, 10, 32, 20, 14})
	fmt.Println(atm.Withdraw(2250))
}

var mp = map[int]int{
	4: 500,
	3: 200,
	2: 100,
	1: 50,
	0: 20,
}

type ATM struct {
	Denom []int
}

// Denom []int{20$,50$,100$,200$,500$}

func Constructor() ATM {
	return ATM{
		Denom: make([]int, 5),
	}
}

func (this *ATM) Deposit(bc []int) {
	for i := 0; i < len(bc); i++ {
		this.Denom[i] += bc[i]
	}
}
func (this *ATM) Withdraw(amount int) []int {
	footprint := make([]int, 5)
	result := make([]int, 5)
	copy(footprint, this.Denom)

	for i := 4; i >= 0; i-- {
		val := mp[i]
		if this.Denom[i] > 0 && amount >= val {
			del := amount / val
			if del > this.Denom[i] {
				del = this.Denom[i]
			}

			amount -= val * del
			result[i] += del
			this.Denom[i] -= del
		}
	}

	if amount != 0 {
		this.Denom = footprint
		return []int{-1}
	}

	return result
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
