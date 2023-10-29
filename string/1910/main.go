package main

import "fmt"

func main() {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(removeOccurrences("axxxxyyyyb", "xy"))
}

func removeOccurrences(s string, part string) string {
	b := []byte(s)

	var temp []byte
	for i := 0; i < len(b); i++ {
		temp = append(temp, b[i])
		if string(temp) == part {
			b = append(b[:i-len(temp)+1], b[i+1:]...)
			i = -1
			temp = nil
		}
		if len(temp) >= len(part) {
			temp = temp[1:]
		}
	}

	return string(b)
}
