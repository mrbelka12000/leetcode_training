package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home//foo/"))
}

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	var result []string

	for i := 0; i < len(arr); i++ {
		if arr[i] == "." || arr[i] == "" {
			continue
		}
		if arr[i] == ".." {
			if len(result) != 0 {
				result = result[:len(result)-1]
			}
			continue
		}
		result = append(result, arr[i])
	}

	return "/" + strings.Join(result, "/")
}
