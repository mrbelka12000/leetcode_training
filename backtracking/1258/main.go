package main

import (
	"sort"
	"strings"
)

func main() {

}

func generateSentences(synonyms [][]string, text string) []string {
	if len(synonyms) == 0 || len(text) == 0 {
		return nil
	}

	m = make(map[string][]string)
	used = make(map[string]bool)
	result = nil

	for _, v := range synonyms {
		for _, q := range m[v[0]] {
			m[v[1]] = append(m[v[1]], q)
			m[q] = append(m[q], v[1])
		}
		for _, q := range m[v[1]] {
			m[v[0]] = append(m[v[0]], q)
			m[q] = append(m[q], v[0])
		}
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}

	runner(strings.Split(text, " "))

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func runner(text []string) {
	tmpWord := strings.Join(text, " ")
	if used[tmpWord] {
		return
	}

	used[tmpWord] = true
	result = append(result, tmpWord)

	for i, word := range text {
		if val, ok := m[word]; ok {
			for _, syn := range val {
				tmp := word
				text[i] = syn
				runner(text)
				text[i] = tmp
			}
		}
	}
}

var used map[string]bool
var m map[string][]string
var result []string
