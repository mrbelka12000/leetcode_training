package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	var result []string

	getPhone(0, digits, "", &result)

	return result
}

var letters = []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func getPhone(ind int, digits, tmp string, result *[]string) {
	if ind == len(digits) {
		if tmp != "" {
			*result = append(*result, tmp)
		}
		return
	}

	num := digits[ind] - '0'
	num--

	for _, v := range letters[num] {
		getPhone(ind+1, digits, tmp+string(v), result)
	}
}
