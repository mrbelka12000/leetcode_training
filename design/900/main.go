package main

func main() {

}

type RLEIterator struct {
	arr []int
}

func Constructor(encoding []int) RLEIterator {

	return RLEIterator{
		arr: encoding,
	}
}

func (this *RLEIterator) Next(next int) int {
	// fmt.Println(this.arr,next)
	for i := 0; i < len(this.arr); i += 2 {
		if this.arr[i] <= 0 {
			continue
		}
		if this.arr[i]-next < 0 {
			next = next - this.arr[i]
			this.arr[i] = 0
			continue
		}
		next = this.arr[i] - next
		this.arr[i] = next

		if next >= 0 {
			return this.arr[i+1]
		}
	}

	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */
