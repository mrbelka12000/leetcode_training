package main

import "fmt"

func main() {
	fmt.Println(getHappyString(4, 9))
}

func getHappyString(n int, k int) string {
	arr := []string{"a", "b", "c"}

	var count int = 1
	return runner(arr, "", k, n, &count)
}

func runner(arr []string, tmp string, k, n int, ind *int) string {
	if *ind == k && len(tmp) == n {
		return tmp
	}

	if len(tmp) == n {
		*ind++
		return ""
	}

	for _, v := range arr {
		if len(tmp) == 0 {
			val := runner(arr, string(v), k, n, ind)
			if val != "" {
				return val
			}
			continue
		}

		if string(tmp[len(tmp)-1]) == v {
			continue
		}

		val := runner(arr, tmp+string(v), k, n, ind)
		if val != "" {
			return val
		}
	}

	return ""
}
