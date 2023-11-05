package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(printVertically("HOW ARE YOU"))
	//fmt.Printf("%#v\n", printVertically("TO BE OR NOT TO BE"))
	fmt.Printf("%#v\n", printVertically("CONTEST IS COMING"))
}

//["HAY"
//,"ORO"
//,"WEU"]
//
//["TBONTB"  "TBONTB"
//,"OEROOE"  "O E R OO E "
//,"   T"]   "   T   "

// ["CIC"
// ,"OSO"
// ,"N M"
// ,"T I"
// ,"E N"
// ,"S G"
// ,"T"]
func printVertically(s string) []string {
	sl := strings.Split(s, " ")
	fmt.Println(sl)

	var (
		maxLen int
	)

	for _, v := range sl {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	result := make([]string, maxLen)

	for i := 0; i < len(sl); i++ {
		var curr int
		for j := 0; j < len(sl[i]); j++ {
			result[j] += string(sl[i][j])
			curr = j
		}

		for k := curr + 1; k < maxLen; k++ {
			result[k] += " "
		}
	}
	for i := 0; i < len(result); i++ {
		result[i] = strings.TrimRight(result[i], " ")
	}

	return result
}
