package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

func capitalizeTitle(t string) string {
	t = strings.ToLower(t)
	spl := strings.Split(t, " ")
	for i := 0; i < len(spl); i++ {
		if len(spl[i]) < 3 {
			spl[i] = strings.ToLower(spl[i])
			continue
		}
		b := []byte(spl[i])
		b[0] = b[0] - 32
		spl[i] = string(b)
	}

	return strings.Join(spl, " ")
}
