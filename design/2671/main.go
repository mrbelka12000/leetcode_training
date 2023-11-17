package main

import "fmt"

func main() {
	ft := Constructor()

	ft.Add(1)
	ft.Add(1)
	ft.Add(2)

	ft.DeleteOne(1)

	fmt.Println(ft.HasFrequency(1))
	fmt.Println(ft.HasFrequency(2))
}

type FrequencyTracker struct {
	mp   map[int]int
	freq map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		mp:   make(map[int]int),
		freq: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	this.mp[number]++

	val := this.mp[number]

	this.freq[val]++
	this.freq[val-1]--

	// fmt.Println(this,number,"add")
}

func (this *FrequencyTracker) DeleteOne(number int) {
	// fmt.Println(this,number,"delete")
	val, ok := this.mp[number]
	if ok {
		this.mp[number]--
	}

	if this.freq[val] > 0 {
		this.freq[val]--
		this.freq[val-1]++
	}

	if val == 1 {
		delete(this.mp, number)
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	// fmt.Println(this, frequency , "has freq")
	val, ok := this.freq[frequency]
	if !ok {
		return false
	}
	if ok && val != 0 {
		return true
	}

	return false
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
