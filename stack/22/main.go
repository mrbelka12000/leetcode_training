package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))

}

func generateParenthesis(num int) []string {

	if num == 0 {
		return []string{""}
	}
	result := []string{}

	l, r := num, num

	gen(l-1, r, "(", &result)

	return result
}

func gen(l, r int, tmp string, result *[]string) {
	if l == 0 && r == 0 {
		*result = append(*result, tmp)
		return
	}

	if l < 0 || r < 0 {
		return
	}

	if l < r {
		gen(l, r-1, tmp+")", result)
	}

	gen(l-1, r, tmp+"(", result)
}
