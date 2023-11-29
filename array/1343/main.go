package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	th := threshold * k

	var result, tmp, firstInd int
	first := arr[0]
	for i := 0; i < len(arr); i++ {
		if i-firstInd == k {
			tmp -= first
			firstInd++
			first = arr[firstInd]
		}

		tmp += arr[i]
		if tmp >= th && i-firstInd == k-1 {
			// fmt.Println(tmp,i)
			result++
		}

		// fmt.Println(first,firstInd, tmp,i)
	}

	return result
}
