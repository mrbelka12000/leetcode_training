package main

func main() {

}

type MagicDictionary struct {
	store []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.store = append(this.store, dictionary...)
}

func (this *MagicDictionary) Search(searchWord string) bool {
	// fmt.Println(searchWord)
	for i := 0; i < len(this.store); i++ {
		if canChange(this.store[i], searchWord) {
			// this.store[i] = "test"
			return true
		}
	}

	return false
}

func canChange(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	var check bool

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if check {
				return false
			}
			check = true
		}
	}

	if !check {
		return false
	}

	return true
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
