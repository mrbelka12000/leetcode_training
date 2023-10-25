package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(41))
}

var nums = []int{1000, 500, 100, 50, 10, 5, 1}

var mp = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

func intToRoman(num int) string {
	var result string
	for num > 0 {
		for i, v := range nums {
			if num >= v {
				str := fmt.Sprint(num)
				//fmt.Println(str)
				if str[0] == '4' {
					result += mp[nums[i]]
					result += mp[nums[i-1]]
					temp := 4
					for temp*10 < nums[i-1] {
						temp *= 10
					}
					num -= temp
					break
				}
				if str[0] == '9' {
					result += mp[nums[i+1]]
					result += mp[nums[i-1]]

					temp := 9
					for temp*10 < nums[i-1] {
						temp *= 10
					}
					num -= temp
					break
				}
				num -= v
				result += mp[v]
				break
			}
		}
	}

	return result
}
