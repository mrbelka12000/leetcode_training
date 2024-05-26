package main

import "fmt"

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
}

func minCost(colors string, neededTime []int) int {
	result = 0

	runner(colors, neededTime, 0)

	return result
}

var result int

func runner(cl string, nt []int, ind int) int {
	if ind >= len(cl)-1 {
		return ind
	}

	i := runner(cl, nt, ind+1)

	if cl[i] == cl[ind] {
		if nt[i] > nt[ind] {
			result += nt[ind]
			return i
		} else {
			result += nt[i]
			return ind
		}
	}

	return ind

}
