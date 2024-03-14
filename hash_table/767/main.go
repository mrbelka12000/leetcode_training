package main

import "fmt"

func main() {
	fmt.Println(reorganizeString("vvvlo"))

}

func reorganizeString(s string) string {
	// mp := make(map[byte]int)
	mp := [26]int{}

	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}

	return runner("", mp)
}

func runner(tmp string, mp [26]int) string {
	var (
		maxVal int
		val    byte
		bb     int
		check  bool
	)

	for k, v := range mp {
		if v == 0 {
			continue
		}

		check = true
		if len(tmp) > 0 {
			if byte(k+'a') == tmp[len(tmp)-1] {
				continue
			}
		}

		if v > maxVal {
			maxVal = v
			bb = k
			val = byte(k + 'a')
		}
	}

	if !check {
		return tmp
	}

	if val == 0 {
		return ""
	}

	mp[bb]--

	return runner(tmp+string(val), mp)
}
