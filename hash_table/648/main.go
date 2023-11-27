package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))

}

func replaceWords(d []string, s string) string {
	arr := strings.Split(s, " ")

	for i := 0; i < len(arr); i++ {
		var root string
		for j := 0; j < len(d); j++ {
			ind := strings.Index(arr[i], d[j])
			if ind == 0 && ind != len(arr[i])-len(d[j]) {
				if root == "" || len(root) > len(d[j]) {
					root = d[j]
				}
			}
		}
		if root != "" {
			arr[i] = root
		}
	}

	fmt.Println(arr)

	return strings.Join(arr, " ")
}
