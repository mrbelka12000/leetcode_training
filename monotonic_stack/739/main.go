package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))

}

func dailyTemperatures(tmp []int) []int {
	res := make([]int, len(tmp))
	var stack []int

	for i, v := range tmp {
		for len(stack) > 0 && tmp[stack[len(stack)-1]] < v {
			ind := stack[len(stack)-1]
			res[ind] = i - ind
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
