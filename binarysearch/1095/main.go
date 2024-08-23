package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 1 }

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := peakIndexInMountainArray(mountainArr)

	if ind := findTargetInLeftArr(mountainArr, target, peak); ind != -1 {
		return ind
	}

	return findTargetInRightArr(mountainArr, target, peak)
}

func findTargetInLeftArr(mountainArr *MountainArray, target, r int) int {
	l := -1
	for r-l > 1 {
		mid := (l + r) / 2
		if mountainArr.get(mid) >= target {
			r = mid
		} else {
			l = mid
		}
	}

	if mountainArr.get(r) != target {
		return -1
	}

	return r
}

func findTargetInRightArr(mountainArr *MountainArray, target, l int) int {
	r := mountainArr.length()
	for r-l > 1 {
		mid := (l + r) / 2
		if mountainArr.get(mid) <= target {
			r = mid
		} else {
			l = mid
		}
	}

	if r >= mountainArr.length() || mountainArr.get(r) != target {
		return -1
	}

	return r
}

func peakIndexInMountainArray(arr *MountainArray) int {
	l, r := 1, arr.length()-2
	for l < r {
		mid := (l + r) / 2
		if arr.get(mid) > arr.get(mid) && arr.get(mid) > arr.get(mid) {
			return mid
		}

		if arr.get(mid) > arr.get(mid+1) {
			r = mid - 1
		}

		if arr.get(mid) > arr.get(mid-1) {
			l = mid + 1
		}
	}

	return r
}
