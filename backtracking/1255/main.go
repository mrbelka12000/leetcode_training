package main

func main() {

}

func maxScoreWords(words []string, letters []byte, score []int) int {
	var result int

	var count [26]int

	for _, v := range letters {
		count[v-'a']++
	}

	runner(words, "", count, score, 0, &result, "")

	return result
}

func runner(words []string, word string, count [26]int, score []int, tmp int, result *int, path string) {

	// fmt.Println(path)
	if word != "" {
		var (
			cnt int
		)

		for i := 0; i < len(word); i++ {
			if count[word[i]-'a'] == 0 {
				// fmt.Println("returned", word, "path->", path)
				return
			}
			count[word[i]-'a']--
			cnt += score[word[i]-'a']
		}

		tmp += cnt
	}

	// fmt.Println("ok", path)

	*result = max(*result, tmp)

	// fmt.Println(words, path)
	// fmt.Println(mp)

	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		tmpWord := words[i]

		words[i] = ""

		runner(words, tmpWord, count, score, tmp, result, path+":"+tmpWord)

		words[i] = tmpWord
	}
}
