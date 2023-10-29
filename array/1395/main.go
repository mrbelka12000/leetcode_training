package main

func main() {

}

func numTeams(r []int) int {
	var result int

	ln := len(r)
	for i := 0; i < ln-2; i++ {
		for j := i + 1; j < ln-1; j++ {
			for k := j + 1; k < ln; k++ {
				if (r[i] < r[j] && r[j] < r[k]) || (r[i] > r[j] && r[j] > r[k]) {
					result++
				}
			}
		}
	}

	return result
}
