package main

import "fmt"

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01", "10"}))
}

func findDifferentBinaryString(nums []string) string {
	dict := make(map[string]bool)
	for _, v := range nums {
		dict[v] = true
	}

	return runner("", dict)
}

func runner(tmp string, dict map[string]bool) string {
	if len(tmp) == len(dict) {
		if dict[tmp] {
			return ""
		}
		return tmp
	}

	val := runner(tmp+"0", dict)
	if val != "" {
		return val
	}

	val = runner(tmp+"1", dict)
	if val != "" {
		return val
	}

	return ""
}
