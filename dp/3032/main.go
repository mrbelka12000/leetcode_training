package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberCount(1, 1000))
}

func numberCount(a int, b int) int {
	arr := getArr(b)

	runner(a, b, 0, arr)

	return arr[len(arr)-1] + 1
}

func runner(a, b, tmp int, arr []int) int {
	if a == b+1 {
		return tmp
	}

	if arr[a] != -1 {
		return arr[a]
	}

	uniq := make(map[int]bool)
	var check bool
	tmpA := a

	for tmpA != 0 {
		if uniq[tmpA%10] {
			check = true
		}
		uniq[tmpA%10] = true
		tmpA /= 10
	}

	if !check {
		tmp++
	}

	arr[a] += runner(a+1, b, tmp, arr)

	return arr[a]
}

func getArr(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
	return arr
}
