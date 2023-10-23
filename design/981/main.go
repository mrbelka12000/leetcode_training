package main

func main() {

}

type TimeMap struct {
	mp map[string][]val
}

type val struct {
	value string
	time  int
}

func newVal(value string, time int) val {
	return val{
		value: value,
		time:  time,
	}
}

func Constructor() TimeMap {
	return TimeMap{
		mp: make(map[string][]val),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	val := newVal(value, timestamp)
	this.mp[key] = append(this.mp[key], val)

}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.mp[key]
	if !ok {
		return ""
	}

	ind := bSearch(values, timestamp)
	if ind == -1 {
		return ""
	}

	return values[ind].value
}

func bSearch(a []val, x int) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid].time == x {
			r = mid // found
			break
		} else if a[mid].time < x {
			start = mid + 1
			r = mid
		} else if a[mid].time > x {
			end = mid - 1
		}
	}

	return r
}
