package main

import "fmt"

func main() {
	//fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 1, 1, 2, 2, 3}))

}

/*
		  *
*	  * * *
* * * * * *
*/
func largestRectangleArea(h []int) int {
	left := findLeftMin(h)
	right := findRightMin(h)
	var result int

	for i := 0; i < len(h); i++ {
		result = max(result, h[i]*(right[i]-left[i]-1))
	}
	return result
}

func findLeftMin(tmp []int) []int {
	res := makeArr(len(tmp), 0)
	res[0] = -1
	for i := 1; i < len(tmp); i++ {
		p := i - 1

		for p >= 0 && tmp[p] >= tmp[i] {
			p = res[p]
		}

		res[i] = p
	}

	return res
}

func findRightMin(tmp []int) []int {
	res := makeArr(len(tmp), 0)
	res[len(tmp)-1] = len(tmp)

	for i := len(tmp) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(tmp) && tmp[p] >= tmp[i] {
			p = res[p]
		}
		res[i] = p
	}
	return res
}

func makeArr(n, val int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = val
	}
	return arr
}
