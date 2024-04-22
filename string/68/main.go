package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))

}

func fullJustify(words []string, maxWidth int) []string {
	var count int

	var tmp, result []string
	for i, v := range words {
		count += len(v)
		tmp = append(tmp, v)
		if i != len(words)-1 {
			if count+len(words[i+1])+len(tmp)-1 >= maxWidth {
				result = append(result, getText(tmp, count, maxWidth))
				count = 0
				tmp = nil
			}
		}
	}

	if len(tmp) == 1 {
		result = append(result, getText(tmp, count, maxWidth))
	} else {
		result = append(result, getLastText(tmp, maxWidth))
	}

	return result
}

func getLastText(str []string, mw int) string {
	result := strings.Builder{}
	var count int
	for i, v := range str {
		count += len(v)
		result.WriteString(v)
		if i != len(str)-1 {
			result.WriteRune(' ')
			count++
		}
	}

	return result.String() + strings.Repeat(" ", mw-count)
}

func getText(str []string, currSize, mw int) string {
	emptySpace := mw - currSize

	var (
		countSpaces    = emptySpace
		countModSpaces int
	)
	if len(str) != 1 {
		countSpaces = emptySpace / (len(str) - 1)
		countModSpaces = emptySpace % (len(str) - 1)
	}

	result := strings.Builder{}
	for pos, v := range str {
		if v == "-1" {
			continue
		}

		result.WriteString(v)
		for i := 0; i < countSpaces; i++ {
			if pos != len(str)-1 || len(str) == 1 {
				result.WriteRune(' ')
			}
		}

		if countModSpaces != 0 {
			countModSpaces--
			result.WriteRune(' ')
		}
	}
	// fmt.Printf("empty=%v, spaces=%v, modSpaces=%v, len(str)=%v\n", emptySpace, countSpaces, countModSpaces, len(str))
	return result.String()
}
