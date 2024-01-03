package main

import "fmt"

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
}

func numberOfBeams(bank []string) int {
	var result int
	var prev int
	for i := 0; i < len(bank); i++ {
		var num int
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				num++
			}
		}
		if prev != 0 && num != 0 {
			result += prev * num
		}

		if num != 0 {
			prev = num
		}
		// fmt.Println(prev,num)
	}

	return result
}
