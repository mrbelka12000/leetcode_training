package main

func main() {

}

func longestPalindrome(s string) int {
	length := 0
	mp := make(map[rune]int)
	for _, v := range s {
		mp[v]++
	}

	var check bool

	for len(mp) > 0 {
		for k, v := range mp {
			if v%2 == 0 {
				length += v
				delete(mp, k)
				continue
			}

			if v == 1 && !check {
				length++
				delete(mp, k)
				check = true
			}

			if v == 1 {
				delete(mp, k)
				continue
			}

			if v > 1 {
				length += 2
				mp[k] -= 2
			}
		}
	}

	return length
}
