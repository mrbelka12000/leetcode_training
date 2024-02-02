package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(sequentialDigits(100, 300))
}

func sequentialDigits(low int, high int) []int {
	var ans []int

	dig := "123456789"

	for i := 0; i < len(dig); i++ {
		runner(low, high, dig[i:], &ans, make(map[int]bool))
	}

	sort.Ints(ans)
	return ans
}

func runner(low, high int, digits string, ans *[]int, mp map[int]bool) {
	// fmt.Println(digits,"-----")

	for i := 0; i < len(digits); i++ {
		intVal, _ := strconv.Atoi(digits[0 : i+1])
		// fmt.Println(intVal)
		if intVal < low || intVal > high || mp[intVal] {
			continue
		}

		*ans = append(*ans, intVal)
		mp[intVal] = true
		// runner(low, high, digits)
	}
}

// func runner(num int)
