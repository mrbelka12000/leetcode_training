package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(addOperators("123456789", 45))
}

func addOperators(num string, target int) []string {
	result = nil
	b := make([]byte, 1, len(num))
	b[0] = num[0]
	runner(num[1:], b, target, string(num[0]))
	return result
}

var result []string

func runner(num string, tmpStr []byte, target int, lastStr string) {
	if len(num) == 0 {
		if getRes(string(tmpStr)) == target {
			result = append(result, string(tmpStr))
		}
		return
	}

	if num[0] == '0' && len(tmpStr) == 0 {
		return
	}

	runner(num[1:], append(tmpStr, []byte{'+', num[0]}...), target, string(num[0]))
	runner(num[1:], append(tmpStr, []byte{'-', num[0]}...), target, string(num[0]))
	runner(num[1:], append(tmpStr, []byte{'*', num[0]}...), target, string(num[0]))
	if len(lastStr) != 0 && lastStr[0] != '0' {
		runner(num[1:], append(tmpStr, []byte{num[0]}...), target, lastStr+string(num[0]))
	}
}

func getRes(str string) int {
	arr := getArr(str)
	if len(arr) == 0 {
		return math.MaxInt32
	}
	for len(arr) != 1 {
		ind := -1
		var check bool
		for i := 0; i < len(arr); i++ {
			if arr[i] == "+" || arr[i] == "-" || arr[i] == "*" {
				if arr[i] == "*" || arr[i] == "/" {
					ind = i
					check = true
					break
				}
				ind = i
			}
		}

		if !check {
			ind = 1
		}

		if ind != -1 {
			num1, _ := strconv.Atoi(arr[ind-1])
			num2, _ := strconv.Atoi(arr[ind+1])
			if arr[ind] == "*" {
				arr[ind-1] = fmt.Sprint(num1 * num2)
			}
			if arr[ind] == "+" {
				arr[ind-1] = fmt.Sprint(num1 + num2)
			}
			if arr[ind] == "-" {
				arr[ind-1] = fmt.Sprint(num1 - num2)
			}
			arr = append(arr[:ind], arr[ind+2:]...)
		}
	}

	val, _ := strconv.Atoi(arr[0])
	return val
}

func getArr(str string) []string {
	var (
		arr   []string
		start int
		check bool
	)

	for i := 0; i < len(str); i++ {
		if str[i] == '+' || str[i] == '-' || str[i] == '*' || str[i] == '/' {
			check = true
			arr = append(arr, str[start:i])
			arr = append(arr, string(str[i]))
			start = i + 1
			continue
		}
	}

	if !check {
		return nil
	}

	arr = append(arr, str[start:len(str)])
	return arr
}
