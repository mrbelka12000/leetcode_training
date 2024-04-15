package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))

}

func letterCasePermutation(str string) []string {
	permutations := []string{str}

	for i := 0; i < len(str); i++ {
		if !unicode.IsLetter(rune(str[i])) {
			continue
		}
		for _, v := range permutations {
			tmp := []rune(v)
			if unicode.IsUpper(tmp[i]) {
				tmp[i] = unicode.ToLower(tmp[i])
			} else {
				tmp[i] = unicode.ToUpper(tmp[i])
			}
			permutations = append(permutations, string(tmp))
		}
	}

	return permutations
}
