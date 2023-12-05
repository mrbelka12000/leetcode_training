package main

func main() {

}

type SnapshotArray struct {
	mp     map[int]map[int]int
	snapID int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		mp: make(map[int]map[int]int),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	if _, ok := this.mp[this.snapID]; !ok {
		this.mp[this.snapID] = make(map[int]int)
	}
	this.mp[this.snapID][index] = val
}

func (this *SnapshotArray) Snap() int {
	defer func() {
		this.snapID++
	}()
	return this.snapID
}

func (this *SnapshotArray) Get(index int, snap_id int) int {

	for i := snap_id; i >= 0; i-- {
		if val, ok := this.mp[i]; ok {
			if v, ok := val[index]; ok {
				return v
			}
		}
	}

	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
