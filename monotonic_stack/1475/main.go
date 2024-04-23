package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}

func finalPrices(prices []int) []int {
	result := make([]int, len(prices))
	copy(result, prices)
	var stack []int

	for i, v := range prices {

		for len(stack) > 0 && prices[stack[len(stack)-1]] >= v {
			ind := stack[len(stack)-1]
			result[ind] -= v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}
