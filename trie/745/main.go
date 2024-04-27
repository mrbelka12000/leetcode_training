package main

func main() {

}

type WordFilter struct {
	store map[string]map[string]int
}

func Constructor(words []string) WordFilter {
	mp := make(map[string]map[string]int)

	for ind, word := range words {
		for i := 0; i < len(word); i++ {
			if mp[word[:i+1]] == nil {
				mp[word[:i+1]] = make(map[string]int)
			}
			for j := 0; j < len(word); j++ {
				mp[word[:i+1]][word[j:]] = ind
			}
		}
	}

	return WordFilter{
		store: mp,
	}
}

func (this *WordFilter) F(pref string, suff string) int {
	val, ok := this.store[pref][suff]
	if !ok {
		return -1
	}

	return val
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
