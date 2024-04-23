package main

func main() {

}

func canSeePersonsCount(h []int) []int {
	result := getArray(len(h))

	var stack []int
	for i := len(h) - 1; i >= 0; i-- {
		count := 0
		for len(stack) > 0 && h[stack[len(stack)-1]] < h[i] {
			count++
			stack = stack[:len(stack)-1]
			// st = st[:len(st)-1]
		}

		if len(stack) != 0 {
			count++
		}

		result[i] = count

		stack = append(stack, i)
		// st = append(st, h[i])
	}

	return result
}

func getArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < len(arr)-1; i++ {
		// arr[i] = 1
	}
	return arr
}
