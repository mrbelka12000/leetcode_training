package main

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
}

func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1
	d, u := 0, len(matrix)-1

	var arr []int

	for l <= r && d <= u {
		for i := u; i >= d; i-- {
			arr = append(arr, matrix[i][l])
		}

		for i := l; i <= r; i++ {
			if i != l {
				arr = append(arr, matrix[d][i])
			}

			if len(arr) != 0 {
				matrix[d][i] = arr[0]
				if i != r {
					arr = arr[1:]
				}
			}
		}

		for i := d; i <= u; i++ {
			if i != d {
				arr = append(arr, matrix[i][r])
			}
			if len(arr) != 0 {
				matrix[i][r] = arr[0]
				if i != u {
					arr = arr[1:]
				}
			}
		}

		for i := r; i >= l; i-- {
			if i != r {
				arr = append(arr, matrix[u][i])
			}
			if len(arr) != 0 {
				matrix[u][i] = arr[0]
				if i != l {
					arr = arr[1:]
				}
			}
		}

		for i := u; i >= d; i-- {
			if i != u {
				arr = append(arr, matrix[i][l])
			}
			if len(arr) != 0 {
				matrix[i][l] = arr[0]
				if i != d {
					arr = arr[1:]
				}
			}
		}

		l++
		r--
		d++
		u--
		arr = nil
	}
}
