package main

import "fmt"

func main() {
	fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
}

func getFolderNames(names []string) []string {
	mp := make(map[string]struct{})
	var result []string

	for _, v := range names {
		file := getFileName(v, mp)
		mp[file] = struct{}{}
		result = append(result, file)
	}

	check = make(map[string]int)
	return result
}

var check = map[string]int{}

func getFileName(file string, mp map[string]struct{}) string {
	// fmt.Println(mp,file)
	_, ok := mp[file]
	if !ok {
		return file
	}

	layout := file

	inc, found := check[layout]
	if !found {
		inc = 1
	}

	for ok {
		file = fmt.Sprintf("%v(%v)", layout, inc)
		_, ok = mp[file]
		inc++
	}

	check[layout] = inc

	return file
}
