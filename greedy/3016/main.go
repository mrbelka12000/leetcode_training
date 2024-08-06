package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
}

func minimumPushes(word string) int {
	var arr [26]int

	for i := 0; i < len(word); i++ {
		arr[word[i]-'a']++
	}

	var store []info

	for i, v := range arr {
		if v == 0 {
			continue
		}
		store = append(store, info{
			char:  byte(i) + 'a',
			count: v,
		})
	}
	sort.Slice(store, func(i, j int) bool {
		return store[i].count > store[j].count
	})

	mp := make(map[byte]int)

	ind := 2
	click := 1
	pos := 0
	for pos != len(store) {
		if ind == 10 {
			ind = 2
			click++
		}

		n := store[pos]

		mp[n.char] = click

		pos++
		ind++
	}

	var result int
	for _, v := range store {
		result += (mp[v.char] * v.count)
	}

	return result
}

type info struct {
	char  byte
	count int
}
