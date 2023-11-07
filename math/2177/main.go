package main

import "fmt"

func main() {
	fmt.Println(sumOfThree(4))
	fmt.Println(sumOfThree(33))
}

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return nil
	}

	result := num / 3

	return []int64{result - 1, result, result + 1}
}
