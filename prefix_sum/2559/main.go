package main

func main() {

}

func vowelStrings(words []string, queries [][]int) []int {
	pref := make([]int, len(words)+1)

	for i, w := range words {
		pref[i+1] = pref[i]
		if isVowel(w) {
			pref[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = pref[q[1]+1] - pref[q[0]]
	}

	return ans
}

func isVowel(s string) bool {
	n := len(s) - 1
	return (s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u') && (s[n] == 'a' || s[n] == 'e' || s[n] == 'i' || s[n] == 'o' || s[n] == 'u')
}

// ["aba","bcb","ece","aa","e"]
//   1      1     2    3    4
//   4      3     3    2    1
