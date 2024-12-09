package main

import (
	"fmt"
)

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
}

func minNumberOfFrogs(croaks string) int {
	var (
		arr    [26][]int
		result = -1
		values = "croak"
	)

	for i := 0; i < len(croaks); i++ {
		arr[croaks[i]-'a'] = append(arr[croaks[i]-'a'], i)

		if isValid(arr) {
			var (
				diff        int
				notPrevFrog bool
			)

			for j := len(values) - 1; j >= 1; j-- {
				diff = arr[values[j]-'a'][0] - arr[values[j-1]-'a'][0]
				if diff != 1 {
					notPrevFrog = true
				}
				arr[values[j]-'a'] = arr[values[j-1]-'a'][1:]
			}
			arr[values[0]-'a'] = arr[values[0]-'a'][1:]

			if result == -1 {
				result = 1
			}
			if notPrevFrog {
				result++
			}
			fmt.Println(arr)
		}
	}

	return result
}

func isValid(arr [26][]int) bool {
	return len(arr['c'-'a']) > 0 &&
		len(arr['r'-'a']) > 0 &&
		len(arr['o'-'a']) > 0 &&
		len(arr['a'-'a']) > 0 &&
		len(arr['k'-'a']) > 0
}
