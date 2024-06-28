package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int {
	return 1
}

func search(reader ArrayReader, target int) int {
	r := getLen(reader)
	l := -1
	var found bool
	for r-l > 1 {
		mid := (l + r) / 2
		val := reader.get(mid)
		if val >= target {
			if val == target {
				found = true
			}
			r = mid
		} else {
			l = mid
		}
	}

	if !found {
		return -1
	}

	return r
}

func getLen(reader ArrayReader) int {
	l, r := 0, 100000

	for r-l > 1 {
		mid := (l + r) / 2
		val := reader.get(mid)
		if val == math.MaxInt32 {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}
