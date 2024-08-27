package main

import (
	"fmt"
)

func main() {
	fmt.Println(findStrobogrammatic(7))
}

func findStrobogrammatic(n int) []string {
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	var result []string

	runner(n, "", "", &result)

	return result
}

func runner(n int, tmp, check string, result *[]string) {
	if n <= 0 {
		if len(tmp+check) > 1 && tmp[0] != '0' {
			*result = append(*result, tmp+check)
		}
		return
	}

	for i := 0; i <= 9; i++ {
		char := "-"
		if i == 8 {
			char = "8"
			if n == 1 {
				runner(n-1, tmp+fmt.Sprint(i), check, result)
			} else {
				runner(n-2, tmp+fmt.Sprint(i), char+check, result)
			}
		} else if i == 6 {
			char = "9"
			if n == 1 {
			} else {
				runner(n-2, tmp+fmt.Sprint(i), char+check, result)
			}
		} else if i == 9 {
			char = "6"
			if n == 1 {
			} else {
				runner(n-2, tmp+fmt.Sprint(i), char+check, result)
			}
		} else if i == 1 {
			char = "1"
			if n == 1 {
				runner(n-1, tmp+fmt.Sprint(i), check, result)
			} else {
				runner(n-2, tmp+fmt.Sprint(i), char+check, result)
			}
		} else if i == 0 {
			char = "0"
			if n == 1 {
				runner(n-1, tmp+fmt.Sprint(i), check, result)
			} else {
				runner(n-2, tmp+fmt.Sprint(i), char+check, result)
			}
		}
	}
}
