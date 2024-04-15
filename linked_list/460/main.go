package main

func main() {

}

type (
	LFUCache struct {
		dic   map[int]*info
		store []*info
		cp    int
	}
	info struct {
		count int
		value int
		pos   int
	}
)

func Constructor(capacity int) LFUCache {
	return LFUCache{
		dic:   make(map[int]*info),
		store: make([]*info, 0),
		cp:    capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if i, ok := this.dic[key]; ok {
		i.count++
		if i.pos != len(this.store)-1 {
			if i.count >= this.store[i.pos+1].count {
				this.store[i.pos+1].pos--
				i.pos++
				i, this.store[i.pos+1] = this.store[i.pos+1], i
			}
		}
		return i.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if i, ok := this.dic[key]; ok {
		this.Get(key)

		i.value = value
		return
	}

	if len(this.store) >= this.cp {
		this.store
	} else {

	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
