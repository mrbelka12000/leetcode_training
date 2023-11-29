package main

import "sort"

func main() {

}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	mp1 := make(map[byte]int)
	mp2 := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		mp1[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		mp2[word2[i]]++
	}

	if len(mp1) != len(mp2) {
		return false
	}

	var arr1, arr2 []int
	for k, v := range mp1 {
		if _, ok := mp2[k]; !ok {
			return false
		}
		arr1 = append(arr1, v)
	}

	for k, v := range mp2 {
		if _, ok := mp1[k]; !ok {
			return false
		}
		arr2 = append(arr2, v)
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	// fmt.Println(string(arr1), string(arr2))

	return true
}
