package main

import "fmt"

func main() {
	//fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(shiftingLetters("bad", []int{10, 20, 30}))
}

func shiftingLetters(s string, shifts []int) string {
	b := []byte(s)

	var sum int

	for _, v := range shifts {
		sum += v
	}

	for i := 0; i < len(b); i++ {
		bb := int(b[i]) + sum
		if bb > 122 {
			num := bb - 97
			count := num / 26
			b[i] = byte(num - count*26 + 97)
		} else {
			b[i] = byte(bb)
		}
		sum -= shifts[i]
	}

	return string(b)
}
